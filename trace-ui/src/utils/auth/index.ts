// @ts-ignore
import Cookies from 'js-cookie'

const TOKEN_KEY = "token"

class TokenUtils {

  getToken() {
    // return Cookies.get(TOKEN_KEY)
    return 'hello'
  }

  setToken(token: string) {
    Cookies.set(TOKEN_KEY, token, {expires: 7})
  }

  removeToken() {
    Cookies.remove(TOKEN_KEY)
  }
}


const tokenUtils = new TokenUtils()

export default tokenUtils