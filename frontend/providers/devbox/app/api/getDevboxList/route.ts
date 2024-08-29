import { NextRequest } from 'next/server'

import { authSession } from '@/services/backend/auth'
import { jsonRes } from '@/services/backend/response'
import { getK8s } from '@/services/backend/kubernetes'
import { devboxKey, publicDomainKey } from '@/constants/devbox'

export const dynamic = 'force-dynamic'

export async function GET(req: NextRequest) {
  try {
    const headerList = req.headers

    const { k8sCustomObjects, namespace, k8sCore } = await getK8s({
      kubeconfig: await authSession(headerList)
    })

    const { body: devboxBody }: any = await k8sCustomObjects.listNamespacedCustomObject(
      'devbox.sealos.io',
      'v1alpha1',
      namespace,
      'devboxes',
      undefined,
      undefined,
      undefined,
      undefined,
      undefined
    )

    const { body: runtimeBody }: any = await k8sCustomObjects.listClusterCustomObject(
      'devbox.sealos.io',
      'v1alpha1',
      'runtimes'
    )

    // 对devboxBody进行处理，增加运行时和网络的一些字段
    const res = devboxBody.items.map(async (item: any) => {
      const devboxName = item.metadata.name
      const runtimeName = item.spec.runtimeRef.name
      const runtime = runtimeBody.items.find((item: any) => item.metadata.name === runtimeName)

      item.spec.runtimeType = runtime?.spec.classRef
      item.spec.runtimeVersion = runtime?.metadata.name

      const { body: ingresses }: any = await k8sCustomObjects.listNamespacedCustomObject(
        'networking.k8s.io',
        'v1',
        namespace,
        'ingresses',
        undefined,
        undefined,
        undefined,
        undefined,
        `${devboxKey}=${devboxName}`
      )
      const { body: certificates }: any = await k8sCustomObjects.listNamespacedCustomObject(
        'cert-manager.io',
        'v1',
        namespace,
        'certificates',
        undefined,
        undefined,
        undefined,
        undefined,
        `${devboxKey}=${devboxName}`
      )
      const customDomain = certificates.items[0]?.spec.dnsNames[0]
      const ingressList = ingresses.items.map((item: any) => {
        return {
          networkName: item.metadata.name,
          port: item.spec.rules[0].http.paths[0].backend.service.port.number,
          protocol: item.metadata.annotations['nginx.ingress.kubernetes.io/backend-protocol'],
          openPublicDomain: !!item.metadata.labels[publicDomainKey],
          publicDomain: item.metadata.labels[publicDomainKey],
          customDomain: customDomain || ''
        }
      })
      const { body: service } = await k8sCore.readNamespacedService(devboxName, namespace)

      item.networks = item.spec.network.extraPorts.map(async (network: any) => {
        const matchingIngress = ingressList.find(
          (ingress: any) => ingress.port === network.containerPort
        )

        const servicePort = service.spec?.ports?.find(
          (port: any) => port.port === network.containerPort
        )
        const servicePortName = servicePort?.name

        if (matchingIngress) {
          return {
            networkName: matchingIngress.networkName,
            port: matchingIngress.port,
            portName: servicePortName,
            protocol: matchingIngress.protocol,
            openPublicDomain: matchingIngress.openPublicDomain,
            publicDomain: matchingIngress.publicDomain,
            customDomain: matchingIngress.customDomain
          }
        }

        return {
          ...network,
          port: network.containerPort
        }
      })
      item.networks = await Promise.all(item.networks)
      return item
    })
    const resp = await Promise.all(res)
    return jsonRes({ data: resp })
  } catch (err: any) {
    return jsonRes({
      code: 500,
      error: err
    })
  }
}
