import {
  Box,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalCloseButton,
  Flex,
  Center,
  Button,
  MenuButton,
  useDisclosure
} from '@chakra-ui/react'
import { useState } from 'react'
import dynamic from 'next/dynamic'
import { SealosMenu } from '@sealos/ui'
import { useTranslations } from 'next-intl'
import { useQuery } from '@tanstack/react-query'

import MyIcon from '@/components/Icon'
import MyTable from '@/components/MyTable'
import { useLoading } from '@/hooks/useLoading'
import { useDevboxStore } from '@/stores/devbox'
import { DevboxVersionListItemType } from '@/types/devbox'

const EditVersionDesModal = dynamic(
  () => import('@/app/[lang]/(platform)/(empty)/components/EditVersionDesModal')
)
const ReleaseModal = dynamic(
  () => import('@/app/[lang]/(platform)/(empty)/components/releaseModal')
)

const Version = ({
  isOpen,
  onClose,
  devboxId
}: {
  isOpen: boolean
  onClose: () => void
  devboxId: string
}) => {
  const t = useTranslations()
  const { Loading } = useLoading()
  const [initialized, setInitialized] = useState(false)
  const [onOpenPublish, setOnOpenPublish] = useState(false)
  const [currentVersion, setCurrentVersion] = useState<DevboxVersionListItemType | null>(null)
  const { devboxVersionList, setDevboxVersionList } = useDevboxStore()
  const { isOpen: isOpenEdit, onOpen: onOpenEdit, onClose: onCloseEdit } = useDisclosure()

  const { refetch } = useQuery(['initDevboxVersionList'], () => setDevboxVersionList(devboxId), {
    onSettled() {
      setInitialized(true)
    }
  })

  const columns: {
    title: string
    dataIndex?: keyof DevboxVersionListItemType
    key: string
    render?: (item: DevboxVersionListItemType) => JSX.Element
    minWidth?: string
  }[] = [
    {
      title: t('version_number'),
      key: 'id',
      render: (item: DevboxVersionListItemType) => (
        <Box color={'grayModern.900'} pl={'12px'}>
          v {item.id}
        </Box>
      )
    },
    {
      title: t('creation_time'),
      dataIndex: 'createTime',
      key: 'createTime'
    },
    {
      title: t('version_description'),
      key: 'description',
      render: (item: DevboxVersionListItemType) => (
        <Flex alignItems="center" className="hover-container">
          <Box
            color={'grayModern.900'}
            overflow={'hidden'}
            textOverflow={'ellipsis'}
            whiteSpace={'nowrap'}
            w={'250px'}>
            {item.description}
          </Box>
          <Box ml={'1px'} className="hover-button" display={'none'}>
            <MyIcon
              name="edit"
              w={'16px'}
              color={'grayModern.600'}
              onClick={() => {
                setCurrentVersion(item)
                onOpenEdit()
              }}
            />
          </Box>
        </Flex>
      ),
      minWidth: '300px'
    },
    {
      title: t('control'),
      key: 'control',
      render: (item: DevboxVersionListItemType) => (
        <Flex>
          <Button
            mr={5}
            height={'32px'}
            w={'50px'}
            size={'sm'}
            fontSize={'base'}
            bg={'grayModern.150'}
            color={'grayModern.900'}
            _hover={{
              color: 'brightBlue.600'
            }}
            onClick={() => {}}>
            {t('online')}
          </Button>
          <SealosMenu
            width={100}
            Button={
              <MenuButton as={Button} variant={'square'} w={'30px'} h={'30px'}>
                <MyIcon name={'more'} color={'grayModern.600'} />
              </MenuButton>
            }
            menuList={[
              {
                child: (
                  <>
                    <MyIcon name={'delete'} w={'16px'} />
                    <Box ml={2}>{t('delete')}</Box>
                  </>
                ),
                menuItemStyle: {
                  _hover: {
                    color: 'red.600',
                    bg: 'rgba(17, 24, 36, 0.05)'
                  }
                },
                onClick: () => {} // TODO: 添加删除逻辑
              }
            ]}
          />
        </Flex>
      )
    }
  ]

  return (
    <Modal isOpen={isOpen} onClose={onClose} lockFocusAcrossFrames={false}>
      <ModalOverlay />
      <ModalContent minW={'900px'}>
        <ModalHeader display={'flex'} alignItems={'center'}>
          <Box flex={1}>{t('version_history')}</Box>
          <ModalCloseButton top={'10px'} right={'10px'} />
        </ModalHeader>
        <ModalBody>
          <Box px={'32px'} minH={'700px'}>
            <Flex h={'90px'} alignItems={'center'}>
              <Center mr={'10px'} height={'46px'}>
                <MyIcon name="pods" w={'20px'} h={'20px'} />
              </Center>
              <Box fontSize={'md'} color={'grayModern.900'} fontWeight={'bold'}>
                {t('version_list')}
              </Box>
              <Box
                w={'2px'}
                h={'24px'}
                bg={'grayModern.200'}
                mx={'16px'}
                my={'8px'}
                borderRadius={'base'}
              />
              <Box ml={'8px'} fontSize={'md'} fontWeight={'bold'} color={'grayModern.500'}>
                ( {devboxVersionList.length} )
              </Box>
              <Box flex={1}></Box>
              <Button
                minW={'100px'}
                h={'35px'}
                variant={'solid'}
                onClick={() => setOnOpenPublish(true)}>
                {t('release_version')}
              </Button>
            </Flex>
            <Loading loading={!initialized} />
            {devboxVersionList.length === 0 && initialized ? (
              <Box>{t('no_versions')}</Box>
            ) : (
              <>
                <MyTable columns={columns} data={devboxVersionList} />
              </>
            )}
          </Box>
        </ModalBody>
      </ModalContent>
      {!!currentVersion && (
        <EditVersionDesModal
          version={currentVersion}
          onSuccess={refetch}
          isOpen={isOpenEdit}
          onClose={onCloseEdit}
        />
      )}
      {!!onOpenPublish && (
        <ReleaseModal
          onSuccess={refetch}
          onClose={() => {
            setOnOpenPublish(false)
          }}
          devboxId={devboxId}
        />
      )}
    </Modal>
  )
}

export default Version
