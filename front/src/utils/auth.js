import Cookies from 'js-cookie'

const TokenKey = '_s'// 'Admin-Token'

export function getToken() {
  const signature = Cookies.get('_s')
  if (signature !== '') {
    return signature
  }

  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}
