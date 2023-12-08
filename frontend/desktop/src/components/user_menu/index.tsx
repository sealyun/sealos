import Account from '@/components/account';
import Notification from '@/components/notification';
import useSessionStore from '@/stores/session';
import {Box, Center, Flex, FlexProps, Image, useDisclosure} from '@chakra-ui/react';
import LangSelectSimple from '../LangSelect/simple';
import Iconfont from '../iconfont';
import GithubComponent from './github';
import {ImageFallBackUrl, useSystemConfigStore} from '@/stores/config';
import {ReactElement, useState} from 'react';
import RegionToggle from '@/components/region/RegionToggle';

enum UserMenuKeys {
    LangSelect,
    Notification,
    Account,
    Region
}

export default function Index(props: { userMenuStyleProps?: FlexProps }) {
    const [notificationAmount, setNotificationAmount] = useState(0);
    const accountDisclosure = useDisclosure();
    const showDisclosure = useDisclosure();

    const {systemConfig} = useSystemConfigStore();
    const userInfo = useSessionStore((state) => state.session);
    const {
        userMenuStyleProps = {
            alignItems: 'center',
            position: 'absolute',
            top: '42px',
            right: '42px',
            cursor: 'pointer',
            gap: '16px'
        }
    } = props;

    const baseItemStyle = {
        w: '36px',
        h: '36px',
        background: 'rgba(244, 246, 248, 0.7)',
        boxShadow: '0px 1.2px 2.3px rgba(0, 0, 0, 0.2)'
    };

    const buttonList: {
        click?: () => void;
        button: ReactElement;
        content: ReactElement;
        key: UserMenuKeys;
    }[] = [
        {
            key: UserMenuKeys.Notification,
            button: (
                <Iconfont iconName="icon-notifications" width={20} height={20} color="#24282C"></Iconfont>
            ),
            click: () => showDisclosure.onOpen(),
            content: (
                <Notification
                    key={'notification'}
                    disclosure={showDisclosure}
                    onAmount={(amount) => setNotificationAmount(amount)}
                />
            )
        },
        {
            key: UserMenuKeys.Account,
            button: (
                <Image
                    width={'36px'}
                    height={'36px'}
                    borderRadius="full"
                    src={userInfo?.user?.avatar || ''}
                    fallbackSrc={ImageFallBackUrl}
                    alt="user avator"
                />
            ),
            click: () => accountDisclosure.onOpen(),
            content: <Account disclosure={accountDisclosure} key={'avatar'}/>
        }
    ];
    return (
        <Flex {...userMenuStyleProps}>
            <RegionToggle/>
            <LangSelectSimple {...baseItemStyle} />
            {systemConfig?.showGithubStar && <GithubComponent {...baseItemStyle} />}
            {buttonList.map((item) => (
                <Flex
                    key={item.key}
                    borderRadius={'50%'}
                    justifyContent={'center'}
                    alignItems={'center'}
                    position={'relative'}
                    {...baseItemStyle}
                >
                    <Center w="100%" h="100%" onClick={item.click} cursor={'pointer'}>
                        {item.button}
                    </Center>
                    {item.content}
                    {item.key === UserMenuKeys.Notification && notificationAmount > 0 && (
                        <Box
                            position={'absolute'}
                            top={0}
                            right={0.5}
                            w={'8px'}
                            h={'8px'}
                            borderRadius={'50%'}
                            backgroundColor={'rgba(255, 132, 146, 1)'}
                        ></Box>
                    )}
                </Flex>
            ))}
        </Flex>
    );
}
