import React, { Dispatch, useCallback, useState } from 'react';
import { Box, Flex, Button, useDisclosure, Center } from '@chakra-ui/react';
import type { DBDetailType } from '@/types/db';
import { useRouter } from 'next/router';
import { restartDB, pauseDBByName, startDBByName } from '@/api/db';
import { useToast } from '@/hooks/useToast';
import { useConfirm } from '@/hooks/useConfirm';
import { defaultDBDetail } from '@/constants/db';
import DBStatusTag from '@/components/DBStatusTag';
import MyIcon from '@/components/Icon';
import dynamic from 'next/dynamic';
import { useTranslation } from 'next-i18next';

const DelModal = dynamic(() => import('./DelModal'));
const BackupModal = dynamic(() => import('./BackupModal'));

const Header = ({
  db = defaultDBDetail,
  isLargeScreen = true,
  setShowSlider
}: {
  db: DBDetailType;
  isLargeScreen: boolean;
  setShowSlider: Dispatch<boolean>;
}) => {
  const { t } = useTranslation();
  const router = useRouter();
  const { toast } = useToast();
  const {
    isOpen: isOpenDelModal,
    onOpen: onOpenDelModal,
    onClose: onCloseDelModal
  } = useDisclosure();
  const { openConfirm: openRestartConfirm, ConfirmChild: RestartConfirmChild } = useConfirm({
    content: t('Confirm Restart')
  });
  const { openConfirm: onOpenPause, ConfirmChild: PauseChild } = useConfirm({
    content: t('Pause Hint')
  });

  const [loading, setLoading] = useState(false);

  const handleRestartApp = useCallback(async () => {
    try {
      setLoading(true);
      await restartDB(db);
      toast({
        title: 'Restart Success',
        status: 'success'
      });
    } catch (error: any) {
      toast({
        title:
          typeof error === 'string'
            ? error
            : error.message || t('Restart Error') || 'Restart Error',
        status: 'error'
      });
      console.error(error);
    }
    setLoading(false);
  }, [db, t, toast]);

  const handlePauseApp = useCallback(async () => {
    try {
      setLoading(true);
      await pauseDBByName(db);
      toast({
        title: t('Pause Success') || 'Pause Success',
        status: 'success'
      });
    } catch (error: any) {
      toast({
        title:
          typeof error === 'string' ? error : error.message || t('Pause Error') || 'Pause Error',
        status: 'error'
      });
      console.error(error);
    }
    setLoading(false);
  }, [db, t, toast]);

  const handleStartApp = useCallback(async () => {
    try {
      setLoading(true);
      await startDBByName(db);
      toast({
        title: t('Start Success') || 'Start Success',
        status: 'success'
      });
    } catch (error: any) {
      toast({
        title:
          typeof error === 'string' ? error : error.message || t('Start Error') || 'Start Error',
        status: 'error'
      });
      console.error(error);
    }
    setLoading(false);
  }, [db, t, toast]);

  return (
    <Flex h={'86px'} alignItems={'center'}>
      <Center width={'36px'} height={'36px'} onClick={() => router.replace('/dbs')}>
        <MyIcon name="arrowLeft" w={'24px'} />
      </Center>
      <Box ml={'4px'} mr={'18px'} fontSize={'2xl'} fontWeight={'bold'} color={'grayModern.900'}>
        {router.query.name || db.dbName}
      </Box>
      <DBStatusTag status={db.status} conditions={db.conditions} showBorder />
      {!isLargeScreen && (
        <Box mx={4}>
          <Button
            minW={'97px'}
            h={'40px'}
            variant={'outline'}
            leftIcon={<MyIcon name="detail" w={'16px'} />}
            onClick={() => setShowSlider(true)}
          >
            {t('Details')}
          </Button>
        </Box>
      )}
      <Box flex={1} />

      {/* btns */}
      {/* Migrate */}
      {/* <Button
        mr={5}
        h={'36px'}
        borderColor={'myGray.200'}
        leftIcon={<MyIcon name={'change'} w={'14px'} />}
        isLoading={loading}
        variant={'base'}
        bg={'white'}
        onClick={() => {
          router.push(`/db/migrate?name=${db.dbName}&dbType=${db.dbType}`);
        }}
      >
        {t('Migrate')}
      </Button> */}
      {db.status.value !== 'Stopped' && (
        <Button
          mr={5}
          minW={'97px'}
          h={'40px'}
          variant={'outline'}
          leftIcon={<MyIcon name={'change'} w={'20px'} />}
          isLoading={loading}
          isDisabled={db.status.value === 'Updating' && !db.isDiskSpaceOverflow}
          onClick={() => {
            router.push(`/db/edit?name=${db.dbName}`);
          }}
        >
          {t('Update')}
        </Button>
      )}
      {db.status.value === 'Stopped' ? (
        <Button
          mr={5}
          minW={'97px'}
          h={'40px'}
          variant={'outline'}
          leftIcon={<MyIcon name="continue" w={'14px'} />}
          isLoading={loading}
          onClick={handleStartApp}
        >
          {t('Continue')}
        </Button>
      ) : (
        <Button
          mr={5}
          minW={'97px'}
          h={'40px'}
          variant={'outline'}
          leftIcon={<MyIcon name="pause" w={'20px'} />}
          isLoading={loading}
          isDisabled={db.status.value === 'Updating'}
          onClick={onOpenPause(handlePauseApp)}
        >
          {t('Pause')}
        </Button>
      )}

      {db.status.value !== 'Stopped' && (
        <Button
          mr={5}
          minW={'97px'}
          h={'40px'}
          variant={'outline'}
          leftIcon={<MyIcon name="restart" w={'20px'} />}
          isDisabled={db.status.value === 'Updating'}
          onClick={openRestartConfirm(handleRestartApp)}
          isLoading={loading}
        >
          {t('Restart')}
        </Button>
      )}

      <Button
        minW={'97px'}
        h={'40px'}
        variant={'outline'}
        leftIcon={<MyIcon name="delete" w={'20px'} />}
        isLoading={loading}
        isDisabled={db.status.value === 'Updating'}
        onClick={onOpenDelModal}
      >
        {t('Delete')}
      </Button>

      {/* modal */}
      <RestartConfirmChild />
      <PauseChild />
      {isOpenDelModal && (
        <DelModal
          dbName={db.dbName}
          onClose={onCloseDelModal}
          onSuccess={() => router.replace('/dbs')}
        />
      )}
    </Flex>
  );
};

export default React.memo(Header);
