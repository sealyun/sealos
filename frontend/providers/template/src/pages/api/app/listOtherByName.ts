import { dbProviderKey, deployManagerKey, templateDeployKey } from '@/constants/keys';
import { authSession } from '@/services/backend/auth';
import { CRDMeta, getK8s } from '@/services/backend/kubernetes';
import { jsonRes } from '@/services/backend/response';
import { ApiResp } from '@/services/kubernet';
import { IncomingMessage } from 'http';
import type { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, res: NextApiResponse<ApiResp>) {
  try {
    const { instanceName } = req.query as { instanceName: string };
    const { namespace, k8sCore, k8sCustomObjects, k8sBatch, k8sAuth } = await getK8s({
      kubeconfig: await authSession(req.headers)
    });
    const labelSelector = `${templateDeployKey}=${instanceName}`;

    const appCRD: CRDMeta = {
      group: 'app.sealos.io',
      version: 'v1',
      namespace: namespace,
      plural: 'apps'
    };

    // secret
    const secretPromise = k8sCore.listNamespacedSecret(
      namespace,
      undefined,
      undefined,
      undefined,
      undefined,
      labelSelector
    );

    // job
    const jobPromise = k8sBatch.listNamespacedJob(
      namespace,
      undefined,
      undefined,
      undefined,
      undefined,
      labelSelector
    );

    // issuer
    const certIssuerPromise = k8sCustomObjects.listNamespacedCustomObject(
      'cert-manager.io',
      'v1',
      namespace,
      'issuers',
      undefined,
      undefined,
      undefined,
      undefined,
      labelSelector
    ) as Promise<{
      response: IncomingMessage;
      body: {
        items: { kind?: string }[];
        kind: 'IssuerList';
      };
    }>;

    // app cr
    const appCrdResourcePromise = k8sCustomObjects.listNamespacedCustomObject(
      appCRD.group,
      appCRD.version,
      appCRD.namespace,
      appCRD.plural,
      undefined,
      undefined,
      undefined,
      undefined,
      labelSelector
    ) as Promise<{
      response: IncomingMessage;
      body: {
        items: { kind?: string }[];
        kind: 'AppList';
      };
    }>;

    // role
    const rolePromise = k8sAuth.listNamespacedRole(
      namespace,
      undefined,
      undefined,
      undefined,
      undefined,
      `${labelSelector},!${dbProviderKey}`
    );
    const roleBindingPromise = k8sAuth.listNamespacedRoleBinding(
      namespace,
      undefined,
      undefined,
      undefined,
      undefined,
      `${labelSelector},!${dbProviderKey}`
    );
    const saPromise = k8sCore.listNamespacedServiceAccount(
      namespace,
      undefined,
      undefined,
      undefined,
      undefined,
      `${labelSelector},!${dbProviderKey}`
    );
    const configMapPromise = k8sCore.listNamespacedConfigMap(
      namespace,
      undefined,
      undefined,
      undefined,
      undefined,
      `${labelSelector},!${dbProviderKey},!${deployManagerKey}`
    );
    // 使用 Promise.allSettled 获取所有结果 [secretResult, jobResult, customResourceResult]
    const result = await Promise.allSettled([
      secretPromise,
      jobPromise,
      appCrdResourcePromise,
      certIssuerPromise,
      rolePromise,
      roleBindingPromise,
      saPromise,
      configMapPromise
    ]);

    const data = result
      .map((res) => {
        if (res.status === 'fulfilled') {
          return res.value.body.items.map((item) => {
            // console.log(item, '==');
            return {
              ...item,
              kind: item.kind ? item.kind : res.value?.body?.kind?.replace('List', '')
            };
          });
        }
      })
      .filter((item) => Array.isArray(item) && item.length > 0);

    jsonRes(res, { data: data });
  } catch (err: any) {
    jsonRes(res, {
      code: 500,
      error: err
    });
  }
}
