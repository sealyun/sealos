import * as k8s from '@kubernetes/client-node';
import type { NextApiRequest, NextApiResponse } from 'next';
import { CRDMeta, K8sApi, ListCRD } from '../../../services/backend/kubernetes';
import { JsonResp } from '../response';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  const config = req.body.kubeconfig;
  const kc = K8sApi(config);
  const kube_user = kc.getCurrentUser();
  if (kube_user === null) {
    return res.status(404);
  }
  const infra_meta: CRDMeta = {
    group: 'infra.sealos.io',
    version: 'v1',
    // namespace: GetUserDefaultNameSpace(kube_user.name),
    namespace: 'infra-system',
    plural: 'infras'
  };

  try {
    const listCrd = await ListCRD(kc, infra_meta);
    JsonResp(listCrd.body, res);
  } catch (err) {
    if (err instanceof k8s.HttpError) {
      console.log(err.body.message);
    }
    JsonResp(err, res);
  }
}
