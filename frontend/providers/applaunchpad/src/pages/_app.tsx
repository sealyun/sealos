import { useEffect } from 'react';
import Head from 'next/head';
import type { AppProps } from 'next/app';
import { ChakraProvider } from '@chakra-ui/react';
import { theme } from '@/constants/theme';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import Router from 'next/router';
import NProgress from 'nprogress'; //nprogress module
import { sealosApp, createSealosApp } from 'sealos-desktop-sdk/app';
import { useConfirm } from '@/hooks/useConfirm';
import throttle from 'lodash/throttle';
import { useGlobalStore } from '@/store/global';
import { useLoading } from '@/hooks/useLoading';
import { getServiceEnv, SEALOS_DOMAIN } from '@/store/static';
import { useRouter } from 'next/router';
import { appWithTranslation } from 'next-i18next';

import 'nprogress/nprogress.css';
import '@/styles/reset.scss';

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

const App = ({ Component, pageProps }: AppProps) => {
  const router = useRouter();
  const { setScreenWidth, loading, setLastRoute } = useGlobalStore();
  const { Loading } = useLoading();
  const { openConfirm, ConfirmChild } = useConfirm({
    title: 'jump_prompt',
    content: 'jump_message'
  });

  useEffect(() => {
    getServiceEnv();
    NProgress.start();
    const response = createSealosApp();

    (async () => {
      try {
        const res = await sealosApp.getSession();
        localStorage.setItem('session', JSON.stringify(res));
        console.log('app init success');
      } catch (err) {
        console.log('App is not running in desktop');
        if (!process.env.NEXT_PUBLIC_MOCK_USER) {
          openConfirm(() => {
            window.open(`https://${SEALOS_DOMAIN}`, '_self');
          })();
        }
      }
    })();
    NProgress.done();

    return response;
  }, [openConfirm]);

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

  // record route
  useEffect(() => {
    return () => {
      setLastRoute(router.asPath);
    };
  }, [router.pathname, setLastRoute]);

  return (
    <>
      <Head>
        <title>Sealos deploy Manager</title>
        <meta name="description" content="Generated by Sealos Team" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <QueryClientProvider client={queryClient}>
        <ChakraProvider theme={theme}>
          <Component {...pageProps} />
          <ConfirmChild />
          <Loading loading={loading} />
        </ChakraProvider>
      </QueryClientProvider>
    </>
  );
};

export default appWithTranslation(App);
