import getCookieValue from "../static/js/getCookieValues";

export default function ({env, $axios}) {
  $axios.interceptors.request.use(config => {
    // クッキーからトークンを取り出してヘッダーに添付する
    config.baseURL = env.apiProtocol + '://' + env.apiHost + ':' + env.apiPort + env.apiBaseRoot;
    config.xsrfCookieName = "_csrf";
    config.xsrfHeaderName = "X-CSRF-Token";
    config.withCredentials = true;
    config.headers = {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
      'X-Requested-With': 'XMLHttpRequest',
      'X-CSRF-Token': getCookieValue('_csrf')
    };
    return config
  });
}
