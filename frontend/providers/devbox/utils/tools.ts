import yaml from 'js-yaml'
import { useMessage } from '@sealos/ui'
import { useTranslations } from 'next-intl'
import * as jsonpatch from 'fast-json-patch'

import { YamlKindEnum } from '@/constants/devbox'
import type { DevboxKindsType, DevboxPatchPropsType } from '@/types/devbox'
import { frameworkVersionMap, languageVersionMap, osVersionMap } from '@/stores/static'

export const cpuFormatToM = (cpu = '0') => {
  if (!cpu || cpu === '0') {
    return 0
  }
  let value = parseFloat(cpu)

  if (/n/gi.test(cpu)) {
    value = value / 1000 / 1000
  } else if (/u/gi.test(cpu)) {
    value = value / 1000
  } else if (/m/gi.test(cpu)) {
    value = value
  } else {
    value = value * 1000
  }
  if (value < 0.1) return 0
  return Number(value.toFixed(4))
}

export const memoryFormatToMi = (memory = '0') => {
  if (!memory || memory === '0') {
    return 0
  }

  let value = parseFloat(memory)

  if (/Ki/gi.test(memory)) {
    value = value / 1024
  } else if (/Mi/gi.test(memory)) {
    value = value
  } else if (/Gi/gi.test(memory)) {
    value = value * 1024
  } else if (/Ti/gi.test(memory)) {
    value = value * 1024 * 1024
  } else {
    console.log('Invalid memory value')
    value = 0
  }

  return Number(value.toFixed(2))
}

export const storageFormatToNum = (storage = '0') => {
  return +`${storage.replace(/gi/i, '')}`
}

export const printMemory = (val: number) => {
  return val >= 1024 ? `${Math.round(val / 1024)} Gi` : `${val} Mi`
}

export function downLoadBlob(content: BlobPart, type: string, fileName: string) {
  const blob = new Blob([content], { type })

  const url = URL.createObjectURL(blob)

  const link = document.createElement('a')
  link.href = url
  link.download = fileName

  link.click()
}

export const obj2Query = (obj: Record<string, string | number>) => {
  let str = ''
  Object.entries(obj).forEach(([key, val]) => {
    if (val) {
      str += `${key}=${val}&`
    }
  })

  return str.slice(0, str.length - 1)
}

export const useCopyData = () => {
  const { message: toast } = useMessage()
  const t = useTranslations()

  return {
    copyData: (data: string, title: string = 'copy_success') => {
      try {
        const textarea = document.createElement('textarea')
        textarea.value = data
        document.body.appendChild(textarea)
        textarea.select()
        document.execCommand('copy')
        document.body.removeChild(textarea)
        toast({
          title: t(title),
          status: 'success',
          duration: 1000
        })
      } catch (error) {
        console.error(error)
        toast({
          title: t('copy_failed'),
          status: 'error'
        })
      }
    }
  }
}

export const str2Num = (str?: string | number) => {
  return !!str ? +str : 0
}

export const getErrText = (err: any, def = '') => {
  const msg: string = typeof err === 'string' ? err : err?.message || def || ''
  msg && console.log('error =>', msg)
  return msg
}

export const getValueDefault = (valueIndex: string) => {
  return (
    languageVersionMap[valueIndex]?.[0]?.id ||
    frameworkVersionMap[valueIndex]?.[0]?.id ||
    osVersionMap[valueIndex]?.[0]?.id ||
    undefined
  )
}

/**
 * patch yamlList and get action
 */
export const patchYamlList = ({
  parsedOldYamlList,
  parsedNewYamlList,
  originalYamlList
}: {
  parsedOldYamlList: string[]
  parsedNewYamlList: string[]
  originalYamlList: DevboxKindsType[]
}) => {
  const oldFormJsonList = parsedOldYamlList
    .map((item) => yaml.loadAll(item))
    .flat() as DevboxKindsType[]

  const newFormJsonList = parsedNewYamlList
    .map((item) => yaml.loadAll(item))
    .flat() as DevboxKindsType[]

  const actions: DevboxPatchPropsType = []

  // find delete
  oldFormJsonList.forEach((oldYamlJson) => {
    const item = newFormJsonList.find(
      (item) => item.kind === oldYamlJson.kind && item.metadata?.name === oldYamlJson.metadata?.name
    )
    if (!item && oldYamlJson.metadata?.name) {
      actions.push({
        type: 'delete',
        kind: oldYamlJson.kind as `${YamlKindEnum}`,
        name: oldYamlJson.metadata?.name
      })
    }
  })

  // find create and patch
  newFormJsonList.forEach((newYamlJson) => {
    const oldFormJson = oldFormJsonList.find(
      (item) =>
        item.kind === newYamlJson.kind && item?.metadata?.name === newYamlJson?.metadata?.name
    )

    if (oldFormJson) {
      const patchRes = jsonpatch.compare(oldFormJson, newYamlJson)

      if (patchRes.length === 0) return

      /* Generate a new json using the formPatchResult and the crJson */
      const actionsJson = (() => {
        try {
          /* find cr json */
          let crOldYamlJson = originalYamlList.find(
            (item) =>
              item.kind === oldFormJson?.kind &&
              item?.metadata?.name === oldFormJson?.metadata?.name
          )

          if (!crOldYamlJson) return newYamlJson
          crOldYamlJson = JSON.parse(JSON.stringify(crOldYamlJson))

          if (!crOldYamlJson) return newYamlJson

          /* generate new json */
          const _patchRes: jsonpatch.Operation[] = patchRes
            .map((item) => {
              let jsonPatchError = jsonpatch.validate([item], crOldYamlJson)
              if (jsonPatchError?.name === 'OPERATION_PATH_UNRESOLVABLE') {
                switch (item.op) {
                  case 'add':
                  case 'replace':
                    return {
                      ...item,
                      op: 'add' as const,
                      value: item.value ?? ''
                    }
                  default:
                    return null
                }
              }
              return item
            })
            .filter((op): op is jsonpatch.Operation => op !== null)

          const patchResYamlJson = jsonpatch.applyPatch(crOldYamlJson, _patchRes, true).newDocument

          // delete invalid field
          // @ts-ignore
          delete patchResYamlJson.status
          patchResYamlJson.metadata = {
            name: patchResYamlJson.metadata?.name,
            namespace: patchResYamlJson.metadata?.namespace,
            labels: patchResYamlJson.metadata?.labels,
            annotations: patchResYamlJson.metadata?.annotations,
            ownerReferences: patchResYamlJson.metadata?.ownerReferences,
            finalizers: patchResYamlJson.metadata?.finalizers
          }

          return patchResYamlJson
        } catch (error) {
          console.error('ACTIONS JSON ERROR:\n', error)
          return newYamlJson
        }
      })()

      if (actionsJson.kind === YamlKindEnum.Service) {
        // @ts-ignore
        const ports = actionsJson?.spec.ports || []
        console.log(ports)

        // @ts-ignore
        if (ports.length > 1 && !ports[0]?.name) {
          // @ts-ignore
          actionsJson.spec.ports[0].name = 'adaptport'
        }
      }

      console.log('patch result:', oldFormJson.metadata?.name, oldFormJson.kind, actionsJson)

      actions.push({
        type: 'patch',
        kind: newYamlJson.kind as `${YamlKindEnum}`,
        value: actionsJson as any
      })
    } else {
      actions.push({
        type: 'create',
        kind: newYamlJson.kind as `${YamlKindEnum}`,
        value: yaml.dump(newYamlJson)
      })
    }
  })

  return actions
}
