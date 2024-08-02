import { valuationMap } from '@/constants/payment';
import { UserQuotaItemType } from '@/pages/api/getQuota';
import request from '@/service/request';
import useEnvStore from '@/stores/env';
import CpuIcon from '../icons/CpuIcon';
import { MemoryIcon } from '../icons/MemoryIcon';
import { StorageIcon } from '../icons/StorageIcon';
import { ApiResp } from '@/types';
import {
  Box,
  Divider,
  Flex,
  Heading,
  HStack,
  Img,
  Stack,
  StackProps,
  Text
} from '@chakra-ui/react';
import { useQuery } from '@tanstack/react-query';
import { useTranslation } from 'react-i18next';
const QuotaPie = dynamic(() => import('../cost_overview/components/quotaPieChart'), { ssr: false });
import dynamic from 'next/dynamic';
export default function Quota(props: StackProps) {
  const { t } = useTranslation();
  const { data } = useQuery(['quota'], () =>
    request<any, ApiResp<{ quota: UserQuotaItemType[] }>>('/api/getQuota')
  );
  const quota = (data?.data?.quota || [])
    .filter((d) => d.type !== 'gpu')
    .map((d) => {
      return {
        ...d,
        title: t(d.type),
        unit: valuationMap.get(d.type)?.unit,
        bg: valuationMap.get(d.type)?.bg
      };
    });
  return (
    <Stack {...props}>
      {quota.map((item) => (
        <HStack key={item.type} gap={'30px'}>
          <QuotaPie data={item} color={item.bg} />
          <Box>
            <HStack>
              {item.type === 'cpu' ? (
                <CpuIcon color={'grayModern.600'} boxSize={'20px'} />
              ) : item.type === 'memory' ? (
                <MemoryIcon color={'grayModern.600'} boxSize={'20px'} />
              ) : item.type === 'storage' ? (
                <StorageIcon color={'grayModern.600'} boxSize={'20px'} />
              ) : (
                <></>
              )}
              <Text fontSize={'16px'} fontWeight="500" color={'grayModern.900'}>
                {t(item.type)}
              </Text>
            </HStack>
            <HStack fontSize={'14px'} gap="10px">
              <Text size={'sm'} color={'grayModern.600'}>
                {' '}
                {t('Used')}: {item.used}
                {item.unit}
              </Text>
              <Divider
                orientation={'vertical'}
                borderColor={'grayModern.600'}
                bgColor={'grayModern.500'}
                h={'10px'}
                borderWidth={'1px'}
              />
              <Text size={'sm'} color={'grayModern.600'}>
                {' '}
                {t('Remain')}: {item.limit - item.used}
                {item.unit}
              </Text>
              <Divider
                orientation={'vertical'}
                borderColor={'grayModern.600'}
                bgColor={'grayModern.500'}
                h={'10px'}
                borderWidth={'1px'}
              />
              <Text size={'sm'} color={'grayModern.600'}>
                {' '}
                {t('Total')}: {item.limit}
                {item.unit}
              </Text>
            </HStack>
          </Box>
        </HStack>
      ))}
    </Stack>
  );
}
