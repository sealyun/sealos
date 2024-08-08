'use client'

import { useEffect } from 'react'
import NProgress from 'nprogress'
import { usePathname } from 'next/navigation'

import { useGlobalStore } from '@/store/global'

import 'nprogress/nprogress.css'

const MainPage = () => {
  const pathname = usePathname()
  const { setLastRoute } = useGlobalStore()

  useEffect(() => {
    return () => {
      setLastRoute(pathname)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [pathname])

  useEffect(() => {
    const handleRouteChangeStart = () => NProgress.start()
    const handleRouteChangeComplete = () => NProgress.done()
    const handleRouteChangeError = () => NProgress.done()

    document.addEventListener('routeChangeStart', handleRouteChangeStart)
    document.addEventListener('routeChangeComplete', handleRouteChangeComplete)
    document.addEventListener('routeChangeError', handleRouteChangeError)

    return () => {
      document.removeEventListener('routeChangeStart', handleRouteChangeStart)
      document.removeEventListener('routeChangeComplete', handleRouteChangeComplete)
      document.removeEventListener('routeChangeError', handleRouteChangeError)
    }
  }, [])

  return <div>test</div>
}

export default MainPage
