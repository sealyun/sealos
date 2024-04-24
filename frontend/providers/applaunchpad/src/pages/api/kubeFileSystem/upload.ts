import { authSession } from '@/services/backend/auth';
import { getK8s } from '@/services/backend/kubernetes';
import { jsonRes } from '@/services/backend/response';
import { ApiResp } from '@/services/kubernet';
import { KubeFileSystem } from '@/utils/kubeFileSystem';
import formidable from 'formidable';
import type { NextApiRequest, NextApiResponse } from 'next';
import { PassThrough } from 'stream';

export default async function handler(req: NextApiRequest, res: NextApiResponse<ApiResp>) {
  try {
    const { namespace, k8sExec } = await getK8s({
      kubeconfig: await authSession(req.headers)
    });

    const kubefs = new KubeFileSystem(k8sExec);
    const { containerName, path, podName } = req.query as {
      containerName: string;
      podName: string;
      path: string;
    };

    let form: any;
    let task = new Promise<string>((resolve) => {
      form = formidable({
        fileWriteStreamHandler: () => {
          const pass = new PassThrough();
          kubefs
            .upload({
              namespace,
              podName,
              containerName,
              path,
              file: pass
            })
            .then((res) => {
              resolve(res);
            })
            .catch((err) => {
              console.log(err, 111);
            });
          return pass;
        }
      });
    });

    await form.parse(req);
    await task;

    jsonRes(res, { data: null });
  } catch (err: any) {
    console.log(err, 'err');
    jsonRes(res, {
      code: 500,
      error: err
    });
  }
}

export const config = {
  api: {
    bodyParser: false
  }
};
