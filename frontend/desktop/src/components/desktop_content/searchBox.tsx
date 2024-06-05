import useAppStore from '@/stores/app';
import { useConfigStore } from '@/stores/config';
import { TApp } from '@/types';
import { Center, Fade, Flex, Image, Input, List, ListItem } from '@chakra-ui/react';
import { SearchIcon } from '@sealos/ui';
import { useTranslation } from 'next-i18next';
import { useRef, useState } from 'react';
import { blurBackgroundStyles } from './index';

export default function SearchBox() {
  const { t, i18n } = useTranslation();
  const logo = useConfigStore().layoutConfig?.logo;
  const { installedApps: apps, runningInfo, openApp, setToHighestLayerById } = useAppStore();
  const inputRef = useRef<HTMLInputElement>(null);
  const [searchTerm, setSearchTerm] = useState('');

  const getAppNames = (app: TApp) => {
    const names = [app.name];
    if (app.i18n) {
      Object.values(app.i18n).forEach((i18nData) => {
        if (i18nData.name) {
          names.push(i18nData.name);
        }
      });
    }
    return names;
  };

  // Filter apps based on search term
  const filteredApps = apps.filter((app) => {
    const appNames = getAppNames(app);
    return appNames.some((name) => name.toLowerCase().includes(searchTerm.toLowerCase()));
  });

  return (
    <Flex
      gap={'8px'}
      flex={1}
      {...blurBackgroundStyles}
      alignItems={'center'}
      onClick={() => {
        inputRef.current?.focus();
      }}
      position={'relative'}
      cursor={'pointer'}
    >
      <SearchIcon ml={'16px'} width={'16px'} height={'16px'} color={'white'} />
      <Input
        mr={'16px'}
        ref={inputRef}
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        w={'full'}
        outline={'none'}
        type="text"
        placeholder={t('Search Apps') || 'Search Apps'}
        bg={'transparent'}
        outlineOffset={''}
        border={'none'}
        _placeholder={{ color: 'white' }}
        _hover={{
          bg: 'transparent'
        }}
        _focus={{
          bg: 'transparent',
          color: 'white'
        }}
      />

      <Fade
        in={searchTerm !== ''}
        style={{
          position: 'absolute',
          top: '100%',
          width: '100%'
        }}
      >
        <List
          mt={2}
          p={'16px'}
          bg={'white'}
          // bg="linear-gradient(0deg, rgba(49, 84, 231, 0.40) 0%, rgba(49, 84, 231, 0.40) 100%), rgba(17, 24, 36, 0.35)"
          borderRadius="xl"
          boxShadow="0px 4px 30px 0px rgba(17, 24, 36, 0.25)"
          zIndex={1}
          width="100%"
          backdropBlur={'blur(250px)'}
        >
          {filteredApps.length > 0 ? (
            filteredApps.map((app) => (
              <ListItem
                key={app.key}
                p={2}
                cursor="pointer"
                _hover={{ bg: 'rgba(255, 255, 255, 0.07)' }}
                onClick={() => {
                  openApp(app);
                  setSearchTerm('');
                }}
                display={'flex'}
                gap={'10px'}
              >
                <Center
                  w="28px"
                  h="28px"
                  borderRadius={'md'}
                  boxShadow={'0px 2px 6px 0px rgba(17, 24, 36, 0.15)'}
                  backgroundColor={'rgba(255, 255, 255, 0.90)'}
                  backdropFilter={'blur(50px)'}
                >
                  <Image
                    width="20px"
                    height="20px"
                    src={app?.icon}
                    fallbackSrc={logo || '/logo.svg'}
                    draggable={false}
                    alt="app logo"
                  />
                </Center>

                {app?.i18n?.[i18n?.language]?.name
                  ? app?.i18n?.[i18n?.language]?.name
                  : t(app?.name)}
              </ListItem>
            ))
          ) : (
            <ListItem p={2}>{t('No Apps Found') || 'No Apps Found'}</ListItem>
          )}
        </List>
      </Fade>
    </Flex>
  );
}
