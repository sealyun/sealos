import { Flex, Text, Box } from '@chakra-ui/react';
import Iconfont from '../iconfont';

export default function Index() {
  return (
    <Flex
      justifyContent={'center'}
      alignItems={'center'}
      w="110px"
      h="42px"
      background={'rgba(21, 37, 57, 0.6)'}
      boxShadow={'0px 1.16667px 2.33333px rgba(0, 0, 0, 0.2)'}
      position={'absolute'}
      bottom={'80px'}
      borderRadius={'8px'}
    >
      <Box pt={'1px'} pr={'6px'}>
        <Iconfont iconName="icon-apps" width={20} height={20} color="#ffffff"></Iconfont>
      </Box>
      <Text color={'#FFFFFF'} fontSize={'14px'} fontWeight={500}>
        更多应用
      </Text>
    </Flex>
  );
}
