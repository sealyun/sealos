import React, { useCallback, useEffect, useRef, useState } from 'react';
import { getPodLogs } from '@/api/db';
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalCloseButton,
  Box,
  useTheme,
  Flex,
  Button,
  MenuButton
} from '@chakra-ui/react';
import { useLoading } from '@/hooks/useLoading';
import { downLoadBold } from '@/utils/tools';
import styles from '../index.module.scss';
import MyMenu from '@/components/Menu';
import { ChevronDownIcon } from '@chakra-ui/icons';
import { streamFetch } from '@/services/streamFetch';
import { default as AnsiUp } from 'ansi_up';
import { useTranslation } from 'next-i18next';

const LogsModal = ({
  dbName,
  podName,
  dbType,
  podAlias,
  pods = [],
  setLogsPodName,
  closeFn
}: {
  dbName: string;
  podName: string;
  dbType: string;
  podAlias: string;
  pods: { alias: string; podName: string }[];
  setLogsPodName: (name: string) => void;
  closeFn: () => void;
}) => {
  const { t } = useTranslation();
  const theme = useTheme();
  const { Loading } = useLoading();
  const [logs, setLogs] = useState('');
  const [isLoading, setIsLoading] = useState(true);
  const LogBox = useRef<HTMLDivElement>(null);
  const ansi_up = useRef(new AnsiUp());

  const watchLogs = useCallback(() => {
    // dbType is empty. pod may has been deleted
    if (!dbType) return closeFn();

    const controller = new AbortController();
    streamFetch({
      url: '/api/pod/getPodLogs',
      data: {
        podName,
        dbType,
        stream: true
      },
      abortSignal: controller,
      firstResponse() {
        setIsLoading(false);
        // scroll bottom
        setTimeout(() => {
          if (!LogBox.current) return;
          LogBox.current.scrollTo({
            top: LogBox.current.scrollHeight
          });
        }, 500);
      },
      onMessage(text) {
        try {
          setLogs((state) => {
            return state + ansi_up.current.ansi_to_html(text);
          });
        } catch (error) {
          console.log(error);
        }

        // scroll bottom
        setTimeout(() => {
          if (!LogBox.current) return;
          const isBottom =
            LogBox.current.scrollTop === 0 ||
            LogBox.current.scrollTop + LogBox.current.clientHeight + 100 >=
              LogBox.current.scrollHeight;

          isBottom &&
            LogBox.current.scrollTo({
              top: LogBox.current.scrollHeight
            });
        }, 100);
      }
    });
    return controller;
  }, [closeFn, dbType, podName]);

  useEffect(() => {
    const controller = watchLogs();
    return () => {
      controller?.abort();
    };
  }, []);

  const exportLogs = useCallback(async () => {
    const allLogs = await getPodLogs({
      dbName,
      podName,
      stream: false
    });
    downLoadBold(allLogs, 'text/plain', 'log.txt');
  }, [dbName, podName]);

  return (
    <Modal isOpen={true} onClose={closeFn} isCentered={true} lockFocusAcrossFrames={false}>
      <ModalOverlay />
      <ModalContent className={styles.logs} display={'flex'} maxW={'90vw'} h={'90vh'} m={0}>
        <Flex p={4} alignItems={'center'}>
          <Box fontSize={'xl'} fontWeight={'bold'}>
            Pod {t('Logs')}
          </Box>
          <Box px={3}>
            <MyMenu
              width={240}
              Button={
                <MenuButton
                  minW={'240px'}
                  h={'32px'}
                  textAlign={'start'}
                  bg={'myWhite.400'}
                  border={theme.borders.base}
                  borderRadius={'md'}
                >
                  <Flex px={4} alignItems={'center'}>
                    <Box flex={1}>{podAlias}</Box>
                    <ChevronDownIcon ml={2} />
                  </Flex>
                </MenuButton>
              }
              menuList={pods.map((item) => ({
                isActive: item.podName === podName,
                child: <Box>{item.alias}</Box>,
                onClick: () => setLogsPodName(item.podName)
              }))}
            />
          </Box>
          <Button size={'sm'} onClick={exportLogs}>
            {t('Export')}
          </Button>
        </Flex>
        <ModalCloseButton top={'9px'} />
        <Box flex={'1 0 0'} h={0} position={'relative'}>
          <Box
            ref={LogBox}
            h={'100%'}
            whiteSpace={'pre'}
            px={4}
            pb={2}
            overflow={'auto'}
            fontWeight={400}
            fontFamily={'SFMono-Regular,Menlo,Monaco,Consolas,monospace'}
            dangerouslySetInnerHTML={{ __html: logs }}
          ></Box>
          <Loading loading={isLoading} fixed={false} />
        </Box>
      </ModalContent>
    </Modal>
  );
};

export default LogsModal;
