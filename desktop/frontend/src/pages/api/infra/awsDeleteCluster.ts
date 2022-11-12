import * as k8s from '@kubernetes/client-node';
import type { NextApiRequest, NextApiResponse } from 'next';
import { CRDMeta, GetUserDefaultNameSpace, K8sApi } from '../../../services/backend/kubernetes';
import { JsonResp } from '../response';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  const { infraName, kubeconfig } = req.body;
  const kc = K8sApi(kubeconfig);
  const kube_user = kc.getCurrentUser();
  if (kube_user === null) {
    res.status(400);
    return;
  }

  const meta: CRDMeta = {
    group: 'cluster.sealos.io',
    version: 'v1',
    namespace: GetUserDefaultNameSpace(kube_user.name),
    plural: 'clusters'
  };

  try {
    const clusterRes = await kc.makeApiClient(k8s.CustomObjectsApi).deleteClusterCustomObject(
      meta.group,
      meta.version,
      meta.plural,
      // infraName // cluster name
      'zjy-3'
    );

    JsonResp(clusterRes, res);
  } catch (err) {
    if (err instanceof k8s.HttpError) {
      console.log(err.body.message);
    }
    JsonResp(err, res);
  }
}
