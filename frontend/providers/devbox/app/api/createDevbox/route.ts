import { NextRequest } from 'next/server'

import { ApiResp } from '@/services/kubernet'
import { DevboxEditType } from '@/types/devbox'
import { jsonRes } from '@/services/backend/response'
import { authSession } from '@/services/backend/auth'
import { getK8s } from '@/services/backend/kubernetes'
import { json2Devbox } from '@/utils/json2Yaml'

export async function POST(req: NextRequest) {
  try {
    //TODO: zod later
    const { devboxForm } = (await req.json()) as { devboxForm: DevboxEditType }

    const { applyYamlList } = await getK8s({
      kubeconfig: await authSession(req)
    })
    const devbox = json2Devbox(devboxForm)
    await applyYamlList([devbox], 'create')

    // TODO: ApiResp的使用不太好，尝试去除
    return jsonRes({
      data: 'success create devbox'
    })
  } catch (err: any) {
    return jsonRes<ApiResp>({
      code: 500,
      error: err
    })
  }
}
