import Layout from '@/components/layout';
import useSessionStore from '@/stores/session';
import { useColorMode } from '@chakra-ui/react';
import { serverSideTranslations } from 'next-i18next/serverSideTranslations';
import { useRouter } from 'next/router';
import { useEffect } from 'react';

const destination = process.env.NEXT_PUBLIC_SERVICE + 'auth/login';

const Home = (props: any) => {
  const router = useRouter();
  const { colorMode, toggleColorMode } = useColorMode();
  const isUserLogin = useSessionStore((s) => s.isUserLogin);

  useEffect(() => {
    colorMode === 'dark' ? toggleColorMode() : null;
  }, [colorMode, toggleColorMode]);

  useEffect(() => {
    const is_login = isUserLogin();
    if (!is_login && router.pathname !== destination && router.asPath !== destination) {
      router.replace(destination);
    }
  }, [router, isUserLogin]);

  return <Layout>{props.children}</Layout>;
};

export async function getStaticProps({ locale }: { locale: any }) {
  return {
    props: {
      ...(await serverSideTranslations(locale))
    }
  };
}
export default Home;
