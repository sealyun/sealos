import dayjs from 'dayjs';

export function formatTime(date: Date) {
  return dayjs(date).format('YYYY-MM-DD HH:mm:ss');
}

function sanitizeMsg(msg: string) {
  return msg.replace(/%[sdj%]/g, '');
}

export function infoLog(msg: string, obj: Record<string, any> = {}) {
  console.log(`[INFO] ${formatTime(new Date())} ${sanitizeMsg(msg)}`, obj);
}

export function errLog(msg: string, error: any) {
  console.log(`[ERROR] ${formatTime(new Date())}`, {
    message: msg,
    stack: error?.stack,
    ...(error?.config && {
      config: {
        headers: error.config.headers,
        url: error.config.url,
        data: error.config.data
      }
    }),
    ...(error?.body && {
      body: error?.body
    }),
    ...(error?.response && {
      response: {
        status: error.response.status,
        statusText: error.response.statusText
      }
    })
  });
}

export function warnLog(msg: string, obj: Record<string, any> = {}) {
  console.log(`[WARN] ${formatTime(new Date())} ${sanitizeMsg(msg)}`, obj);
}
