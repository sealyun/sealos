/* eslint-disable @next/next/no-img-element */

import React, { useState } from 'react';
import { Spinner } from '@fluentui/react-components';
import clsx from 'clsx';
import { useQuery } from '@tanstack/react-query';
import useAppStore from 'stores/app';
import styles from './index.module.scss';
import DownLoadPage from './download_page';
import DetailPage from './detail_page';
import FrontPage from './front_page';

export default function AppStore() {
  const [page, setPage] = useState(1);
  const { allApps: apps, getAllApps } = useAppStore(({ allApps, getAllApps }) => ({
    allApps,
    getAllApps
  }));
  const { isLoading } = useQuery(['allApps'], () => getAllApps());
  const [opapp, setOpapp] = useState(apps[0]);

  if (isLoading) {
    return <Spinner size={'large'} />;
  }
  const toTab = (x: string) => {
    if (x) {
      setPage(0);
      setTimeout(() => {
        var target = document.getElementById(x);
        if (target) {
          const tsof = target!.parentNode!.parentNode!.scrollTop,
            trof = target.offsetTop;

          if (Math.abs(tsof - trof) > window.innerHeight * 0.1) {
            target!.parentNode!.parentNode!.scrollTop = target.offsetTop;
          }
        }
      }, 200);
    }
  };

  const action = ({ page, appIdentifier }: { page: string; appIdentifier?: string }) => {
    switch (page) {
      case 'page1':
        setPage(1);
        break;
      case 'page2':
        const app = apps.find((item) => item.data.url === appIdentifier);
        if (app) {
          setPage(2);
          setOpapp(app);
        }
        break;
      default:
        break;
    }
  };
  const frontScroll = (e: React.MouseEvent<HTMLElement>) => {
    if (page == 0) {
      let tabs = ['sthome', 'apprib', 'gamerib', 'movrib'],
        mntab = 'sthome',
        mndis = window.innerHeight;

      tabs.forEach((x) => {
        const target = document.getElementById(x);
        if (target) {
          const tsof = target!.parentNode!.parentNode!.scrollTop;
          const trof = target.offsetTop;

          if (Math.abs(tsof - trof) < mndis) {
            mntab = x;
            mndis = Math.abs(tsof - trof);
          }
        }
      });
    }
  };

  return (
    <div className={clsx(styles.appWrap, ' flex h-full ')}>
      <div className={clsx(styles.storeNav, 'mt-4 h-full w-24 flex flex-col')}>
        <a
          href="#"
          className={styles.uicon}
          onClick={(e) => {
            e?.preventDefault();
            action({ page: 'page1' });
          }}
        >
          计算
        </a>
        <a
          href="#"
          className={styles.uicon}
          onClick={(e) => {
            e?.preventDefault();
            action({ page: 'page1' });
          }}
        >
          存储
        </a>
        <a
          href="#"
          className={styles.uicon}
          onClick={(e) => {
            e?.preventDefault();
            action({ page: 'page1' });
          }}
        >
          数据
        </a>
      </div>
      <div className={clsx(styles.restWindow, styles.pageScroll, styles.pageWrapper)}>
        {page === 0 ? <FrontPage /> : null}
        {page === 1 ? <DownLoadPage action={action} apps={(apps.length && apps) || apps} /> : null}
        {page === 2 ? <DetailPage app={opapp} /> : null}
      </div>
    </div>
  );
}
