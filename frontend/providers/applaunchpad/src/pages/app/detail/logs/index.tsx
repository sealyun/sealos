import { useTranslation } from 'next-i18next';
import { useQuery } from '@tanstack/react-query';
import { Box, useTheme, Flex, Divider } from '@chakra-ui/react';

import { useAppStore } from '@/store/app';
import { useToast } from '@/hooks/useToast';
import { serviceSideProps } from '@/utils/i18n';
import DetailLayout from '@/components/layouts/DetailLayout';

import { Header } from './components/Header';
import { Filter } from './components/Filter';
import { LogCounts } from './components/LogCounts';

export default function LogsPage({ appName }: { appName: string }) {
  const { toast } = useToast();
  const { t } = useTranslation();
  const { appDetail } = useAppStore();

  const theme = useTheme();

  const { data: monitorData } = useQuery(
    ['monitor-data', appName],
    async () => {
      return [];
    },
    {
      onError(err) {
        toast({
          title: String(err),
          status: 'error'
        });
      }
    }
  );

  return (
    <DetailLayout appName={appName}>
      <Box flex={1} bg="white" borderRadius="lg" p={4}>
        <>
          <Flex
            mb={4}
            bg={'white'}
            gap={'12px'}
            flexDir={'column'}
            border={theme.borders.base}
            borderRadius={'lg'}
          >
            <Header />
            <Divider />
            <Filter />
          </Flex>
          <Box
            mb={4}
            p={4}
            bg={'white'}
            border={theme.borders.base}
            borderRadius={'lg'}
            flexShrink={0}
            minH={'257px'}
          >
            <LogCounts />
          </Box>
          <Box
            bg={'white'}
            p={4}
            border={theme.borders.base}
            borderRadius={'lg'}
            h={0}
            flex={1}
            minH={'257px'}
          >
            日志数量
          </Box>
        </>
      </Box>
    </DetailLayout>
  );
}

export async function getServerSideProps(content: any) {
  const appName = content?.query?.name || '';

  return {
    props: {
      appName,
      ...(await serviceSideProps(content))
    }
  };
}
