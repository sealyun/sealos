import { verifyAccessToken } from '@/services/backend/auth';
import { jsonRes } from '@/services/backend/response';
import { updateOrder } from '@/services/db/workorder';
import { WorkOrderDB } from '@/types/workorder';
import { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  try {
    const { userId } = await verifyAccessToken(req);

    const { orderId, updates } = req.body as {
      updates: Partial<WorkOrderDB>;
      orderId: string;
    };

    const result = await updateOrder({
      orderId,
      userId: userId,
      updates
    });

    return jsonRes(res, {
      data: result
    });
  } catch (error) {
    jsonRes(res, { code: 500, data: error });
  }
}
