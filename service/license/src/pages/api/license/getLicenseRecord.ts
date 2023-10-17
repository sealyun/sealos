import { authSession } from '@/services/backend/auth';
import { getLicenseRecordsByUid } from '@/services/backend/db/license';
import { jsonRes } from '@/services/backend/response';
import type { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, resp: NextApiResponse) {
  try {
    const payload = await authSession(req.headers);
    if (!payload) return jsonRes(resp, { code: 401, message: 'token verify error' });

    const { page = 1, pageSize = 10 } = req.body as {
      page: number;
      pageSize: number;
    };

    const result = await getLicenseRecordsByUid({
      uid: payload.uid,
      page: page,
      pageSize: pageSize
    });

    return jsonRes(resp, {
      data: result
    });
  } catch (error) {
    jsonRes(resp, { code: 500, data: error });
  }
}
