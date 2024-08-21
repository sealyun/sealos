import { headers } from 'next/headers'
import type { NextRequest } from 'next/server'

import { getK8s } from '@/services/backend/kubernetes'
import { jsonRes } from '@/services/backend/response'
import { authSession } from '@/services/backend/auth'

export async function GET(req: NextRequest) {
  try {
    const headerList = headers()

    const { getUserQuota } = await getK8s({
      kubeconfig: await authSession(headerList)
    })

    const quota = await getUserQuota()

    return jsonRes({
      data: {
        quota
      }
    })
  } catch (error) {
    console.log(error)
    return jsonRes({ code: 500, message: 'get price error' })
  }
}
