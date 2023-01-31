import { useState, useCallback, useMemo, useEffect, useRef } from 'react';
import AppStore from 'applications/app_store';
import Infra from 'applications/infra';
import PgSql from 'applications/pgsql';

import clsx from 'clsx';
import { APPTYPE } from 'constants/app_type';
import useAppStore, { TApp } from 'stores/app';
import AppIcon from '../app_icon';
import AppWindow from '../app_window';
import IframeApp from './iframe_app';
import styles from './index.module.scss';

export default function DesktopContent() {
  const { installedApps: apps, openedApps, openApp, updateAppOrder } = useAppStore((state) => state)
  
  /* icon orders */
  const itemsLen = 18*8 // x:18, y:8
  const gridItems = useMemo(() => new Array(itemsLen).fill(null).map((_, i) => {
    const app = apps.find(item => item.order === i)
    return !!app ? {...app} : null
  }),[apps, itemsLen])
  /* dragging icon */
  const [downingItemIndex, setDowningItemIndex] = useState<number>()

  const isBrowser = typeof window !== 'undefined';
  const desktopWidth = isBrowser ? document.getElementById('desktop')?.offsetWidth || 0 : 0;
  const desktopHeight = isBrowser ? document.getElementById('desktop')?.offsetHeight || 0 : 0;

  function renderApp(appItem: TApp) {
    switch (appItem.type) {
      case APPTYPE.APP:
        if (appItem.name === 'sealos cloud provider') {
          return <Infra />;
        }
        if (appItem.name === 'Postgres') {
          return <PgSql />;
        }
        return <AppStore />;

      case APPTYPE.IFRAME:
        return <IframeApp appItem={appItem} />;

      default:
        break;
    }
  }

  const onDrop = useCallback((e:any, i:number) => {
    setDowningItemIndex(undefined)
    const dom:Element  = e.target
    /* if it doesnot contain "item", it drop in a appGrid */
    if(!dom.classList.contains('item')) return

    if(!downingItemIndex || gridItems[downingItemIndex] === null) return

    // @ts-ignore nextline
    updateAppOrder(gridItems[downingItemIndex], i)
  },[downingItemIndex, gridItems, updateAppOrder])

  return (
    <div className={styles.desktop}>
      {/* 已安装的应用 */}
      <div className={styles.desktopCont}>
        {gridItems.map((item, i:number) => {
          return (
            <div
              key={i}
              className={`item ${styles.dskItem}`}
              draggable={i === downingItemIndex}
              onMouseDown={() => item && setDowningItemIndex(i)}
              onDragOver={(e) => e.preventDefault()}
              onDrop={(e) => onDrop(e, i)}
            >
              {
                !!item ? (
                  <div
                    className={styles.dskApp}
                    onClick={() => {
                      openApp(item)
                    }}
                  >
                    <div className={`${styles.dskIcon}`}>
                      <AppIcon className={clsx('prtclk')} src={item.icon} width="100%" />
                    </div>
                    <div className={styles.appName}>{item.name}</div>
                  </div>
                ) : null
              }
            </div>
          );
        })}
      </div>

      {/* 打开的应用窗口 */}
      {openedApps.map((appItem) => {
        return (
          <AppWindow
            key={appItem.name}
            style={{ height: '100vh' }}
            app={appItem}
            desktopWidth={desktopWidth}
            desktopHeight={desktopHeight}
          >
            {renderApp(appItem)}
          </AppWindow>
        );
      })}
    </div>
  );
}
