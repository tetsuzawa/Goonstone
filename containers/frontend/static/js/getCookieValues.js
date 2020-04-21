/**
 * クッキーの値を取得する
 * @param {String} key 検索するキー
 * @returns {String} キーに対応する値
 */
export default function getCookieValue (key) {
  if (typeof key === 'undefined') {
    return ''
  }
  let val = '';
  document.cookie.split(';').forEach(cookie => {
    const [key, value] = cookie.split('=');
    if (key === key) {
      return val = value
    }
  });

  return val
}
