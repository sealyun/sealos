import yaml from 'js-yaml'
import { NextRequest } from 'next/server'

import { ApiResp } from '@/services/kubernet'
import { DevboxEditType } from '@/types/devbox'
import { jsonRes } from '@/services/backend/response'
import { authSession } from '@/services/backend/auth'
import { getK8s } from '@/services/backend/kubernetes'
import { json2Devbox, json2Ingress, json2Service } from '@/utils/json2Yaml'

export const dynamic = 'force-dynamic'

export async function POST(req: NextRequest) {
  try {
    //TODO: zod later
    const { devboxForm, isEdit } = (await req.json()) as {
      devboxForm: DevboxEditType
      isEdit: boolean
    }
    const headerList = req.headers

    const { applyYamlList, k8sCustomObjects } = await getK8s({
      kubeconfig: await authSession(headerList)
    })
    const devbox = json2Devbox(devboxForm)
    const service = json2Service(devboxForm)
    const ingress = json2Ingress(devboxForm)
    const jsonDevbox = yaml.load(devbox) as object

    if (isEdit) {
      await k8sCustomObjects.patchNamespacedCustomObject(
        'devbox.sealos.io',
        'v1alpha1',
        'default', // TODO: namespace动态获取
        'devboxes',
        devboxForm.name,
        jsonDevbox,
        undefined,
        undefined,
        undefined,
        {
          headers: {
            'Content-Type': 'application/merge-patch+json'
          }
        }
      )
      await applyYamlList([service, ingress], 'update')

      return jsonRes({
        data: 'success update devbox'
      })
    }

    await applyYamlList([devbox, service, ingress], 'create')

    // TODO: ApiResp的使用不太好，尝试去除
    return jsonRes({
      data: 'success create devbox'
    })
  } catch (err: any) {
    return jsonRes<ApiResp>({
      code: 500,
      error: err
    })
  }
}
