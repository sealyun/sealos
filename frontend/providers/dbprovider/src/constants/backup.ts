import { BackupStatusMapType } from '@/types/db';

export enum BackupStatusEnum {
  Completed = 'Completed',
  InProgress = 'InProgress',
  Failed = 'Failed',
  UnKnow = 'UnKnow',
  Running = 'Running',
  Deleting = 'Deleting'
}

export const backupStatusMap: Record<`${BackupStatusEnum}`, BackupStatusMapType> = {
  [BackupStatusEnum.Completed]: {
    label: 'Backup Completed',
    value: BackupStatusEnum.Completed,
    color: '#13B2A9'
  },
  [BackupStatusEnum.InProgress]: {
    label: 'Backup Processing',
    value: BackupStatusEnum.InProgress,
    color: '#CE9629'
  },
  [BackupStatusEnum.Failed]: {
    label: 'Backup Failed',
    value: BackupStatusEnum.Failed,
    color: '#FF5B6E'
  },
  [BackupStatusEnum.Running]: {
    label: 'Backup Running',
    value: BackupStatusEnum.Running,
    color: '#13B2A9'
  },
  [BackupStatusEnum.Deleting]: {
    label: 'Backup Deleting',
    value: BackupStatusEnum.Deleting,
    color: '#FF5B6E'
  },
  [BackupStatusEnum.UnKnow]: {
    label: 'UnKnow',
    value: BackupStatusEnum.UnKnow,
    color: '#787A90'
  }
};

export const BACKUP_TYPE_LABEL_KEY = 'backup-type';
export const BACKUP_REMARK_LABEL_KEY = 'backup-remark';
export const BACKUP_LABEL_KEY = 'kubeblocks.io/restore-from-backup';
export const BACKUP_REPO_DEFAULT_KEY = 'dataprotection.kubeblocks.io/is-default-repo';

export enum BackupTypeEnum {
  manual = 'manual',
  auto = 'auto',
  'UnKnow' = 'UnKnow'
}

export const backupTypeMap: Record<`${BackupTypeEnum}`, { label: string }> = {
  [BackupTypeEnum.manual]: {
    label: 'Manual'
  },
  [BackupTypeEnum.auto]: {
    label: 'Auto'
  },
  [BackupTypeEnum.UnKnow]: {
    label: 'Unknown'
  }
};
