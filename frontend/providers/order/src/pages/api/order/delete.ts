import { authSession } from '@/services/backend/auth';
import { getK8s } from '@/services/backend/kubernetes';
import { jsonRes } from '@/services/backend/response';
import { deleteOrder } from '@/services/db/order';
import { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  try {
    const { namespace } = await getK8s({
      kubeconfig: await authSession(req)
    });

    const { orderID } = req.query as {
      orderID: string;
    };

    const result = await deleteOrder({
      orderID: orderID,
      userID: namespace
    });

    return jsonRes(res, {
      data: result
    });
  } catch (error) {
    jsonRes(res, { code: 500, data: error });
  }
}
