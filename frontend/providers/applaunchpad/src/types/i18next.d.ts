import 'i18next';

import common from '../../public/locales/zh/common.json';

export interface I18nNamespaces {
  common: typeof common;
}

export type I18nNsType = (keyof I18nNamespaces)[];

// declare module 'i18next' {
//   interface CustomTypeOptions {
//     returnNull: false;
//     defaultNS: 'common';
//     resources: I18nNamespaces;
//   }
// }
