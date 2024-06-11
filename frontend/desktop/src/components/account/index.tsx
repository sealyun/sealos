import { useCopyData } from '@/hooks/useCopyData';
import request from '@/services/request';
import useAppStore from '@/stores/app';
import { useConfigStore } from '@/stores/config';
import useSessionStore from '@/stores/session';
import { ApiResp } from '@/types';
import download from '@/utils/downloadFIle';
import { Box, Center, Flex, IconButton, Image, Text } from '@chakra-ui/react';
import {
  CopyIcon,
  DocsIcon,
  DownloadIcon,
  LogoutIcon,
  NotificationIcon,
  SettingIcon
} from '@sealos/ui';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { useTranslation } from 'next-i18next';
import { useRouter } from 'next/router';
import { useMemo, useState } from 'react';
import LangSelectSimple from '../LangSelect/simple';
import { blurBackgroundStyles } from '../desktop_content';
import RegionToggle from '../region/RegionToggle';
import WorkspaceToggle from '../team/WorkspaceToggle';
import PasswordModify from './PasswordModify';
import GithubComponent from './github';

const baseItemStyle = {
  w: '52px',
  h: '40px',
  background: 'rgba(255, 255, 255, 0.07)',
  color: 'white',
  borderRadius: '100px',
  _hover: {
    background: 'rgba(255, 255, 255, 0.15)'
  }
};

export default function Account() {
  const { layoutConfig } = useConfigStore();
  const [showId, setShowId] = useState(true);
  const passwordEnabled = useConfigStore().authConfig?.idp?.password?.enabled;
  const rechargeEnabled = useConfigStore().commonConfig?.rechargeEnabled;
  const logo = useConfigStore().layoutConfig?.logo;
  const router = useRouter();
  const { copyData } = useCopyData();
  const openApp = useAppStore((s) => s.openApp);
  const installApp = useAppStore((s) => s.installedApps);
  const { t } = useTranslation();
  const { delSession, session, setToken } = useSessionStore();
  const user = session?.user;
  const { data } = useQuery({
    queryKey: ['getAmount', { userId: user?.userCrUid }],
    queryFn: () =>
      request<any, ApiResp<{ balance: number; deductionBalance: number }>>(
        '/api/account/getAmount'
      ),
    enabled: !!user
  });
  const balance = useMemo(() => {
    let real_balance = data?.data?.balance || 0;
    if (data?.data?.deductionBalance) {
      real_balance -= data?.data.deductionBalance;
    }
    return real_balance;
  }, [data]);
  const queryclient = useQueryClient();
  const kubeconfig = session?.kubeconfig || '';
  const logout = (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault();
    delSession();
    queryclient.clear();
    router.replace('/signin');
    setToken('');
  };

  return (
    <Box position={'relative'} flex={1}>
      <Flex position={'relative'} zIndex={3} px={'16px'} pt={'20px'} flexDirection={'column'}>
        <Flex alignItems={'center'}>
          <Center width={'36px'} height={'36px'} bg={'white'} borderRadius="full" mr={'10px'}>
            <Image
              width={'24px'}
              height={'24px'}
              borderRadius="full"
              src={user?.avatar || ''}
              fallbackSrc={logo}
              alt="user avator"
            />
          </Center>
          <Box>
            <Text lineHeight={'20px'} color={'white'} fontSize={'14px'} fontWeight={500}>
              {user?.name}
            </Text>
            <Flex
              cursor={'pointer'}
              gap="2px"
              fontSize={'12px'}
              lineHeight={'16px'}
              fontWeight={'500'}
              color={'rgba(255, 255, 255, 0.70)'}
              alignItems={'center'}
            >
              <Text onClick={() => setShowId((s) => !s)}>
                {showId ? `ID: ${user?.userId}` : `NS: ${user?.nsid}`}
              </Text>
              <CopyIcon
                onClick={() => {
                  if (user?.userId && user.nsid) copyData(showId ? user?.userId : user?.nsid);
                }}
                boxSize={'12px'}
                fill={'rgba(255, 255, 255, 0.70)'}
              />
            </Flex>
          </Box>
          <Center
            p={'4px'}
            h={'fit-content'}
            borderRadius={'4px'}
            ml={'auto'}
            cursor={'pointer'}
            _hover={{
              background: 'rgba(255, 255, 255, 0.15)'
            }}
          >
            <LogoutIcon boxSize={'14px'} fill={'white'} />
            <Text ml="4px" color={'white'} fontSize={'12px'} fontWeight={500} onClick={logout}>
              {t('Log Out')}
            </Text>
          </Center>
        </Flex>
        <Flex mt={'16px'} justifyContent={'space-between'}>
          <Center cursor={'pointer'} {...baseItemStyle}>
            <NotificationIcon />
          </Center>
          <LangSelectSimple {...baseItemStyle} />
          {layoutConfig?.common.githubStarEnabled && <GithubComponent {...baseItemStyle} />}
          <Center cursor={'pointer'} {...baseItemStyle}>
            <DocsIcon />
          </Center>
        </Flex>

        <RegionToggle />
        <WorkspaceToggle />
        <Flex
          borderBottom={'1px solid rgba(255, 255, 255, 0.05)'}
          color={'white'}
          fontSize={'base'}
          fontWeight={'bold'}
          justifyContent={'space-between'}
          alignItems={'center'}
          py={'12px'}
          px={'16px'}
        >
          <Text>{t('Account Settings')}</Text>
          <IconButton
            variant={'white-bg-icon'}
            p="4px"
            onClick={() => kubeconfig && copyData(kubeconfig)}
            icon={<SettingIcon boxSize={'16px'} fill={'rgba(255, 255, 255, 0.7)'} />}
            aria-label={'setting'}
          />
        </Flex>
        <Flex
          borderBottom={'1px solid rgba(255, 255, 255, 0.05)'}
          color={'white'}
          fontSize={'base'}
          fontWeight={'bold'}
          justifyContent={'space-between'}
          alignItems={'center'}
          py={'12px'}
          px={'16px'}
        >
          <Text>kubeconfig</Text>
          <Flex alignItems={'center'}>
            <IconButton
              variant={'white-bg-icon'}
              p="4px"
              ml="auto"
              mr="4px"
              onClick={() => kubeconfig && download('kubeconfig.yaml', kubeconfig)}
              icon={
                <DownloadIcon
                  boxSize={'16px'}
                  color={'rgba(255, 255, 255, 0.7)'}
                  fill={'rgba(255, 255, 255, 0.7)'}
                />
              }
              aria-label={'Download kc'}
            />
            <IconButton
              variant={'white-bg-icon'}
              p="4px"
              onClick={() => kubeconfig && copyData(kubeconfig)}
              icon={<CopyIcon boxSize={'16px'} fill={'rgba(255, 255, 255, 0.7)'} />}
              aria-label={'copy kc'}
            />
          </Flex>
        </Flex>
        {passwordEnabled && (
          <Flex
            color={'white'}
            fontSize={'base'}
            fontWeight={'bold'}
            justifyContent={'space-between'}
            alignItems={'center'}
            py={'12px'}
            px={'16px'}
          >
            <Text>{t('changePassword')}</Text>
            <PasswordModify mr="0" />
          </Flex>
        )}
      </Flex>
      <Box
        id="blur-background"
        zIndex={0}
        position={'absolute'}
        top={0}
        left={0}
        w={'full'}
        h={'full'}
        overflow={'hidden'}
        {...blurBackgroundStyles}
      ></Box>
    </Box>
  );
}
