import React, { useCallback, useState } from 'react';
import { useRouter } from 'next/router';
import { Box, Button, Flex, MenuButton } from '@chakra-ui/react';

import DBStatusTag from '@/components/DBStatusTag';
import MyIcon from '@/components/Icon';
import { useTheme } from '@chakra-ui/react';
import { useGlobalStore } from '@/store/global';
import { useToast } from '@/hooks/useToast';
import { restartDB, pauseDBByName, startDBByName } from '@/api/db';
import MyTable from '@/components/Table';
import dynamic from 'next/dynamic';
import MyMenu from '@/components/Menu';
import { useConfirm } from '@/hooks/useConfirm';
import { DBStatusEnum, DBComponentNameMap } from '@/constants/db';
import { printMemory } from '@/utils/tools';
import { useTranslation } from 'next-i18next';
import { CronJobListItemType } from '@/types/job';

const DelModal = dynamic(() => import('@/pages/job/detail/components/DelModal'));

const DBList = ({
  list = [],
  refetchApps
}: {
  list: CronJobListItemType[];
  refetchApps: () => void;
}) => {
  const { t } = useTranslation();
  const { setLoading } = useGlobalStore();
  const { toast } = useToast();
  const theme = useTheme();
  const router = useRouter();
  const [delAppName, setDelAppName] = useState('');

  const { openConfirm: onOpenPause, ConfirmChild: PauseChild } = useConfirm({
    content: t('Pause Hint')
  });

  const handleRestartApp = useCallback(
    async (db: CronJobListItemType) => {
      try {
        setLoading(true);
        await restartDB({ dbName: db.name, dbType: db.dbType });
        toast({
          title: t('Restart Success'),
          status: 'success'
        });
      } catch (error: any) {
        toast({
          title: typeof error === 'string' ? error : error.message || t('Restart Success'),
          status: 'error'
        });
        console.error(error, '==');
      }
      setLoading(false);
    },
    [setLoading, t, toast]
  );

  const handlePauseApp = useCallback(
    async (db: CronJobListItemType) => {
      try {
        setLoading(true);
        await pauseDBByName({ dbName: db.name, dbType: db.dbType });
        toast({
          title: t('Pause Success'),
          status: 'success'
        });
      } catch (error: any) {
        toast({
          title: typeof error === 'string' ? error : error.message || t('Pause Error'),
          status: 'error'
        });
        console.error(error);
      }
      setLoading(false);
      refetchApps();
    },
    [refetchApps, setLoading, t, toast]
  );

  const handleStartApp = useCallback(
    async (db: CronJobListItemType) => {
      try {
        setLoading(true);
        await startDBByName({ dbName: db.name, dbType: db.dbType });
        toast({
          title: t('Start Success'),
          status: 'success'
        });
      } catch (error: any) {
        toast({
          title: typeof error === 'string' ? error : error.message || t('Start Error'),
          status: 'error'
        });
        console.error(error);
      }
      setLoading(false);
      refetchApps();
    },
    [refetchApps, setLoading, t, toast]
  );

  const columns: {
    title: string;
    dataIndex?: keyof CronJobListItemType;
    key: string;
    render?: (item: CronJobListItemType) => JSX.Element;
  }[] = [
    {
      title: 'Name',
      key: 'name',
      render: (item: CronJobListItemType) => {
        return (
          <Box pl={4} color={'myGray.900'} fontSize={'md'}>
            {item.name}
          </Box>
        );
      }
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status'
    },
    {
      title: 'Schedule',
      dataIndex: 'schedule',
      key: 'schedule'
    },
    {
      title: 'Creation Time',
      dataIndex: 'createTime',
      key: 'createTime'
    },
    {
      title: 'Last Schedule Time',
      dataIndex: 'lastScheduleTime',
      key: 'lastScheduleTime'
    },
    {
      title: 'Last Successful Time',
      dataIndex: 'lastSuccessfulTime',
      key: 'lastSuccessfulTime'
    },
    {
      title: 'Operation',
      key: 'control',
      render: (item: CronJobListItemType) => (
        <Flex>
          <Button
            mr={5}
            variant={'base'}
            leftIcon={<MyIcon name={'detail'} transform={'translateY(-1px)'} />}
            px={3}
            onClick={() => {}}
          >
            {t('Details')}
          </Button>
          <MyMenu
            width={100}
            Button={
              <MenuButton
                w={'32px'}
                h={'32px'}
                borderRadius={'sm'}
                _hover={{
                  bg: 'myWhite.400',
                  color: 'hover.iconBlue'
                }}
              >
                <MyIcon name={'more'} px={3} />
              </MenuButton>
            }
            menuList={[
              {
                child: (
                  <>
                    <MyIcon name={'delete'} w={'12px'} />
                    <Box ml={2}>{t('Delete')}</Box>
                  </>
                ),
                onClick: () => setDelAppName(item.name)
              }
            ]}
          />
        </Flex>
      )
    }
  ];

  return (
    <Box bg={'#F3F4F5'} px={'34px'} minH="100vh">
      <Flex h={'88px'} alignItems={'center'}>
        <Box mr={4} p={2} backgroundColor={'#FEFEFE'} border={theme.borders.sm} borderRadius={'sm'}>
          <MyIcon name="logo" w={'24px'} h={'24px'} />
        </Box>
        <Box fontSize={'2xl'} color={'black'}>
          {t('job.list')}
        </Box>
        <Box ml={3} color={'gray.500'}>
          ( {list.length} )
        </Box>
        <Box flex={1}></Box>

        <Button
          flex={'0 0 155px'}
          h={'40px'}
          colorScheme={'primary'}
          leftIcon={<MyIcon name={'plus'} w={'12px'} />}
          variant={'primary'}
          onClick={() => router.push('/job/edit')}
        >
          {t('job.create')}
        </Button>
      </Flex>
      <MyTable columns={columns} data={list} />
      <PauseChild />
      {!!delAppName && (
        <DelModal dbName={delAppName} onClose={() => setDelAppName('')} onSuccess={refetchApps} />
      )}
    </Box>
  );
};

export default DBList;
