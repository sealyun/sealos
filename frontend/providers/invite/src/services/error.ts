export const ERROR_TEXT: Record<string, string> = {
  ETIMEDOUT: '服务器超时'
};

export enum ERROR_ENUM {
  unAuthorization = 'unAuthorization'
}
export const ERROR_RESPONSE: Record<
  any,
  {
    code: number;
    statusText: string;
    message: string;
    data?: any;
  }
> = {
  [ERROR_ENUM.unAuthorization]: {
    code: 403,
    statusText: ERROR_ENUM.unAuthorization,
    message: '凭证错误'
  }
};
