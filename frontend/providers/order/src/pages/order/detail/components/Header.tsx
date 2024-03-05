import MyIcon from '@/components/Icon';
import { useGlobalStore } from '@/store/global';
import { Box, Button, Flex } from '@chakra-ui/react';
import { useTranslation } from 'next-i18next';
import { useRouter } from 'next/router';
import { Dispatch } from 'react';
import AppStatusTag from '@/components/AppStatusTag';
import { OrderDB } from '@/types/order';

const Header = ({
  app,
  isLargeScreen = true,
  setShowSlider
}: {
  app: OrderDB;
  isLargeScreen: boolean;
  setShowSlider: Dispatch<boolean>;
}) => {
  const { t } = useTranslation();
  const router = useRouter();
  const { lastRoute } = useGlobalStore();

  return (
    <Flex w={'100%'} h={'86px'} alignItems={'center'}>
      <Flex alignItems={'center'} cursor={'pointer'} onClick={() => router.replace(lastRoute)}>
        <MyIcon name="arrowLeft" />
        <Box ml={6} fontWeight={'bold'} color={'black'} fontSize={'3xl'} mr="16px">
          {t('Order Detail')}
        </Box>
        {app?.status && <AppStatusTag status={app?.status} showBorder={false} />}
      </Flex>
      <Box flex={1}></Box>
      {!isLargeScreen && (
        <Box mx={4}>
          <Button
            flex={1}
            h={'40px'}
            borderColor={'myGray.200'}
            leftIcon={<MyIcon name="detail" w="16px" h="16px" />}
            variant={'base'}
            bg={'white'}
            onClick={() => setShowSlider(true)}
          >
            {t('Details')}
          </Button>
        </Box>
      )}
      <Button
        _focusVisible={{ boxShadow: '' }}
        mr={5}
        h={'40px'}
        borderColor={'myGray.200'}
        leftIcon={<MyIcon name={'close'} w={'16px'} />}
        variant={'base'}
        bg={'white'}
      >
        {t('Close')}
      </Button>
      <Button
        h={'40px'}
        borderColor={'myGray.200'}
        leftIcon={<MyIcon name="delete" w={'16px'} />}
        variant={'base'}
        bg={'white'}
        _hover={{
          color: '#FF324A'
        }}
      >
        {t('Delete')}
      </Button>
    </Flex>
  );
};

export default Header;
