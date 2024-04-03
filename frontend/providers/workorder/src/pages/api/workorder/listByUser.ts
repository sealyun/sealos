import { verifyAccessToken } from '@/services/backend/auth';
import { jsonRes } from '@/services/backend/response';
import { getAllOrdersByUserId } from '@/services/db/workorder';
import { WorkOrderStatus, WorkOrderType } from '@/types/workorder';
import { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  try {
    const {
      page = 1,
      pageSize = 10,
      orderStatus,
      orderType,
      startTime,
      endTime
    } = req.body as {
      page: number;
      pageSize: number;
      orderType?: WorkOrderType;
      orderStatus?: WorkOrderStatus;
      startTime?: Date;
      endTime?: Date;
    };

    const { userId } = await verifyAccessToken(req);

    let result = await getAllOrdersByUserId({
      userId: userId,
      page: page,
      pageSize: pageSize,
      orderStatus,
      orderType,
      startTime,
      endTime
    });

    return jsonRes(res, {
      data: result
    });
  } catch (error) {
    console.log(error);
    jsonRes(res, { code: 500, error: error });
  }
}
