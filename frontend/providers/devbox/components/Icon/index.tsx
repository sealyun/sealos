import React from 'react'
import { Icon } from '@chakra-ui/react'
import type { IconProps } from '@chakra-ui/react'

const map = {
  codeServer: require('./icons/codeServer.svg').default,
  version: require('./icons/version.svg').default,
  terminal: require('./icons/terminal.svg').default,

  more: require('./icons/more.svg').default,
  podList: require('./icons/podList.svg').default,
  arrowLeft: require('./icons/arrowLeft.svg').default,
  plus: require('./icons/plus.svg').default,
  delete: require('./icons/delete.svg').default,
  restart: require('./icons/restart.svg').default,
  start: require('./icons/start.svg').default,
  pause: require('./icons/pause.svg').default,
  warningInfo: require('./icons/warningInfo.svg').default,
  detail: require('./icons/detail.svg').default,
  logo: require('./icons/logo.svg').default,
  change: require('./icons/change.svg').default,
  formInfo: require('./icons/formInfo.svg').default,
  settings: require('./icons/settings.svg').default,
  copy: require('./icons/copy.svg').default,
  continue: require('./icons/continue.svg').default,
  noEvents: require('./icons/noEvents.svg').default,
  network: require('./icons/network.svg').default,
  warning: require('./icons/warning.svg').default,
  analyze: require('./icons/analyze.svg').default,
  log: require('./icons/log.svg').default,
  statusDetail: require('./icons/statusDetail.svg').default,
  read: require('./icons/read.svg').default,
  unread: require('./icons/unread.svg').default,
  connection: require('./icons/connection.svg').default,
  info: require('./icons/info.svg').default,
  restore: require('./icons/restore.svg').default,
  download: require('./icons/download.svg').default,
  loading: require('./icons/loading.svg').default,
  success: require('./icons/success.svg').default,
  error: require('./icons/error.svg').default,
  currency: require('./icons/currency.svg').default,
  infoCircle: require('./icons/infoCircle.svg').default,
  upperRight: require('./icons/upperRight.svg').default,
  arrowUp: require('./icons/arrowUp.svg').default,
  search: require('./icons/search.svg').default,
  edit: require('./icons/edit.svg').default,
  book: require('./icons/book.svg').default,
  export: require('./icons/export.svg').default,
  pods: require('./icons/pods.svg').default,
  upload: require('./icons/upload.svg').default,
  target: require('./icons/target.svg').default,
  arrowDown: require('./icons/arrowDown.svg').default,
  docs: require('./icons/docs.svg').default,
  vscode: require('./icons/vscode.svg').default,
  monitor: require('./icons/monitor.svg').default,
  response: require('./icons/response.svg').default,
  link: require('./icons/link.svg').default,
  list: require('./icons/list.svg').default,
  maximize: require('./icons/maximize.svg').default,
  chevronDown: require('./icons/chevronDown.svg').default,
  vscodeInsiders: require('./icons/vscodeInsiders.svg').default,
  cursor: require('./icons/cursor.svg').default,
  check: require('./icons/check.svg').default,
  empty: require('./icons/empty.svg').default,
  shutdown: require('./icons/shutdown.svg').default,
  windsurf: require('./icons/windsurf.svg').default,
  rocket: require('./icons/rocket.svg').default,
  jetbrains: require('./icons/jetbrains.svg').default
}

const MyIcon = ({
  name,
  w = 'auto',
  h = 'auto',
  ...props
}: { name: keyof typeof map } & IconProps) => {
  return map[name] ? (
    <Icon as={map[name]} verticalAlign={'text-top'} fill={'currentColor'} w={w} h={h} {...props} />
  ) : null
}

export default MyIcon
