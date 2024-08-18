import {
  Box,
  Button,
  Flex,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Textarea
} from '@chakra-ui/react'
import { useMessage } from '@sealos/ui'
import { useTranslations } from 'next-intl'
import { useCallback, useState } from 'react'

import { releaseDevbox } from '@/api/devbox'
import { DevboxListItemType } from '@/types/devbox'

const ReleaseModal = ({
  onClose,
  onSuccess,
  devbox
}: {
  devbox: DevboxListItemType
  onClose: () => void
  onSuccess: () => void
}) => {
  const t = useTranslations()
  const [tag, setTag] = useState('')
  const { message: toast } = useMessage()
  const [loading, setLoading] = useState(false)
  const [releaseDes, setReleaseDes] = useState('')

  const handleReleaseDevbox = useCallback(async () => {
    if (devbox.status.value === 'Running') {
      toast({
        title: t('devbox_running_cannot_release'),
        status: 'error'
      })
      return
    }

    try {
      setLoading(true)
      await releaseDevbox({
        devboxName: devbox.name,
        tag,
        releaseDes
      })
      toast({
        title: t('release_successful'),
        status: 'success'
      })
      onSuccess()
      onClose()
    } catch (error: any) {
      toast({
        title: typeof error === 'string' ? error : error.message || t('release_failed'),
        status: 'error'
      })
      console.error(error)
    }
    setLoading(false)
  }, [devbox.name, tag, releaseDes, toast, t, onSuccess, onClose])

  return (
    <Modal isOpen onClose={onClose} lockFocusAcrossFrames={false}>
      <ModalOverlay />
      <ModalContent minW={'500px'} mt={'200px'} minH={'300px'}>
        <ModalHeader>
          <Flex alignItems={'center'} gap={'10px'}>
            {t('release_version')}
          </Flex>
        </ModalHeader>
        <ModalCloseButton top={'10px'} right={'10px'} />
        <ModalBody pb={4}>
          <Flex alignItems={'start'} gap={'5px'} mb={'24px'}>
            <Box w={'100px'} fontWeight={'bold'} fontSize={'lg'}>
              {t('image_name')}
              {/* TODO: 获取镜像名 */}
            </Box>
            <Input defaultValue={'test'} />
          </Flex>
          <Flex alignItems={'start'} gap={'5px'}>
            <Box w={'100px'} fontWeight={'bold'} fontSize={'lg'}>
              {t('version_config')}
            </Box>
            <Flex gap={'5px'} direction={'column'}>
              <Box w={'100px'}>{t('version_number')}</Box>
              <Input
                placeholder={t('enter_version_number')}
                value={tag}
                onChange={(e) => setTag(e.target.value)}
                mb={'16px'}
              />
              <Box w={'100px'}>{t('version_description')}</Box>
              <Textarea
                value={releaseDes}
                minH={'150px'}
                onChange={(e) => setReleaseDes(e.target.value)}
                placeholder={t('enter_version_description')}
              />
            </Flex>
          </Flex>
        </ModalBody>
        <ModalFooter>
          {/* TODO: 发版弹窗忘搞了 */}
          <Button
            variant={'solid'}
            onClick={handleReleaseDevbox}
            mr={'20px'}
            width={'60px'}
            isLoading={loading}>
            {t('release')}
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  )
}

export default ReleaseModal
