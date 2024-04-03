import { theme } from '@/constants/theme';
import { useConfirm } from '@/hooks/useConfirm';
import useEnvStore from '@/store/env';
import { useGlobalStore } from '@/store/global';
import useSessionStore from '@/store/session';
import '@/styles/reset.scss';
import { getLangStore, setLangStore } from '@/utils/cookieUtils';
import { ChakraProvider } from '@chakra-ui/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { throttle } from 'lodash';
import { appWithTranslation, useTranslation } from 'next-i18next';
import type { AppProps } from 'next/app';
import Head from 'next/head';
import Router, { useRouter } from 'next/router';
import NProgress from 'nprogress'; //nprogress module
import 'nprogress/nprogress.css';
import { useEffect, useState } from 'react';
import 'react-day-picker/dist/style.css';
import { EVENT_NAME } from 'sealos-desktop-sdk';
import { createSealosApp, sealosApp } from 'sealos-desktop-sdk/app';

//Binding events.
Router.events.on('routeChangeStart', () => NProgress.start());
Router.events.on('routeChangeComplete', () => NProgress.done());
Router.events.on('routeChangeError', () => NProgress.done());

// Create a client
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: false,
      cacheTime: 0
    }
  }
});

function App({ Component, pageProps }: AppProps) {
  const router = useRouter();
  const { initSystemEnv } = useEnvStore();
  const { authUser, delAppSession } = useSessionStore();
  const { i18n } = useTranslation();
  const { setScreenWidth, setLastRoute } = useGlobalStore();
  const [refresh, setRefresh] = useState(false);

  const { openConfirm, ConfirmChild } = useConfirm({
    title: '跳转提示',
    content: '该应用不允许单独使用，点击确认前往 Sealos Desktop 使用。'
  });

  useEffect(() => {
    const response = createSealosApp();

    (async () => {
      const { domain } = await initSystemEnv();
      try {
        const res = await sealosApp.getSession();
        if (!res?.token) return;
        authUser(res.token);
      } catch (err) {
        if (!process.env.NEXT_PUBLIC_MOCK_USER) {
          delAppSession();
          openConfirm(() => {
            window.open(`https://${domain}`, '_self');
          })();
        } else {
          authUser(process.env.NEXT_PUBLIC_MOCK_USER);
        }
      }
    })();
    return response;
  }, []);

  // add resize event
  useEffect(() => {
    const resize = throttle((e: Event) => {
      const documentWidth = document.documentElement.clientWidth || document.body.clientWidth;
      setScreenWidth(documentWidth);
    }, 200);
    window.addEventListener('resize', resize);
    const documentWidth = document.documentElement.clientWidth || document.body.clientWidth;
    setScreenWidth(documentWidth);

    return () => {
      window.removeEventListener('resize', resize);
    };
  }, [setScreenWidth]);

  // init
  useEffect(() => {
    const changeI18n = async (data: any) => {
      const lastLang = getLangStore();
      const newLang = data.currentLanguage;
      if (lastLang !== newLang) {
        i18n?.changeLanguage?.(newLang);
        setLangStore(newLang);
        setRefresh((state) => !state);
      }
    };

    (async () => {
      try {
        const lang = await sealosApp.getLanguage();
        changeI18n({
          currentLanguage: lang.lng
        });
      } catch (error) {
        changeI18n({
          currentLanguage: 'zh'
        });
      }
    })();

    return sealosApp?.addAppEventListen(EVENT_NAME.CHANGE_I18N, changeI18n);
  }, []);

  // record route
  useEffect(() => {
    return () => {
      setLastRoute(router.asPath);
    };
  }, [router.pathname]);

  useEffect(() => {
    const lang = getLangStore() || 'zh';
    i18n?.changeLanguage?.(lang);
  }, [refresh, router.asPath]);

  return (
    <>
      <Head>
        <title>Sealos Work Order</title>
        <meta name="description" content="Generated by Sealos Team" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <QueryClientProvider client={queryClient}>
        <ChakraProvider theme={theme}>
          <Component {...pageProps} />
          <ConfirmChild />
        </ChakraProvider>
      </QueryClientProvider>
    </>
  );
}

export default appWithTranslation(App);
