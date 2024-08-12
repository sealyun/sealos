import JSZip from 'jszip'
import dayjs from 'dayjs'
import React, { useCallback } from 'react'
import { useRouter } from 'next/navigation'
import { Box, Flex, Button } from '@chakra-ui/react'

import MyIcon from '@/components/Icon'
import { downLoadBlob } from '@/utils/tools'
import { useGlobalStore } from '@/stores/global'
import type { YamlItemType } from '@/types/index'

const Header = ({ yamlList, applyCb }: { yamlList: YamlItemType[]; applyCb: () => void }) => {
  const router = useRouter()
  const { lastRoute } = useGlobalStore()

  const handleExportYaml = useCallback(async () => {
    const zip = new JSZip()
    yamlList.forEach((item) => {
      zip.file(item.filename, item.value)
    })
    const res = await zip.generateAsync({ type: 'blob' })
    downLoadBlob(res, 'application/zip', `yaml${dayjs().format('YYYYMMDDHHmmss')}.zip`)
  }, [yamlList])

  return (
    <Flex w={'100%'} px={10} h={'86px'} alignItems={'center'}>
      <Flex alignItems={'center'} cursor={'pointer'} onClick={() => router.replace(lastRoute)}>
        <MyIcon name="arrowLeft" width={'24px'} height={'24px'} />
        <Box fontWeight={'bold'} color={'grayModern.900'} fontSize={'2xl'}>
          {'项目创建'}
        </Box>
      </Flex>
      <Box flex={1}></Box>
      <Button h={'40px'} flex={'0 0 114px'} mr={5} variant={'outline'} onClick={handleExportYaml}>
        {'导出YAML'}
      </Button>
      <Button flex={'0 0 114px'} h={'40px'} variant={'solid'} onClick={applyCb}>
        {'创建'}
      </Button>
    </Flex>
  )
}

export default Header
