import { authSession } from '@/services/backend/auth';
import { createPaymentRecord } from '@/services/backend/db/payment';
import { jsonRes } from '@/services/backend/response';
import { getSealosPay } from '@/services/pay';
import { PaymentDB, PaymentData, PaymentParams, PaymentStatus } from '@/types';
import { formatMoney } from '@/utils/tools';
import type { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  try {
    const { amount, currency, payMethod } = req.body as PaymentParams;

    const userInfo = await authSession(req.headers);
    if (!userInfo) return jsonRes(res, { code: 401, message: 'token verify error' });
    const { sealosPayUrl, sealosPayID, sealosPayKey } = getSealosPay();
    if (!sealosPayUrl)
      return jsonRes(res, { code: 500, message: 'sealos payment has not been activated' });

    const result: PaymentData = await fetch(`${sealosPayUrl}/v1alpha1/pay/session`, {
      method: 'POST',
      body: JSON.stringify({
        appID: 45141910007488120,
        sign: '076f82f8e996d7',
        amount: amount,
        currency: currency,
        user: userInfo.uid,
        payMethod: payMethod
      })
    }).then((res) => res.json());

    let payRecord: PaymentDB = {
      ...result,
      uid: userInfo.uid,
      status: PaymentStatus.PaymentNotPaid,
      amount: formatMoney(parseInt(result.amount)),
      createdAt: new Date(),
      updatedAt: new Date(),
      payMethod: payMethod
    };

    await createPaymentRecord(payRecord);

    return jsonRes(res, {
      data: result
    });
  } catch (error) {
    console.error(error);
    jsonRes(res, { code: 500, data: error });
  }
}
