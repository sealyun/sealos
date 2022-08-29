import { Background } from '@/components/background';
import styles from './index.module.scss';
import Taskbar from '@/components/taskbar';
import DesktopContent from '@/components/desktop_content';
import useAppStore from 'stores/app';

export default function Layout({ children }: any) {
  return (
    <div className={styles.desktopContainer}>
      <Background />
      <DesktopContent />
      <Taskbar />
    </div>
  );
}
