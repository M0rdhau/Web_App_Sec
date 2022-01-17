import login from '../services/login'
import encryptionService from '../services/encryptionsService';

const USER_NAME = 'APPSEC_USERNAME'
const USER_TOKEN = 'APPSEC_TOKEN'

const userReducer = (state = null, action) => {
  switch(action.type){
    case 'LOGIN':
      return action.data
    case 'LOGOUT':
      return null
    default:
      return state
  }
}

export const loginUser = (username, password) => {
  return async dispatch => {
    const data = await login(username, password)
    encryptionService.setToken(data.token)
    window.localStorage.setItem(USER_NAME, data.username)
    window.localStorage.setItem(USER_TOKEN, data.token)
    dispatch({
      type: 'LOGIN',
      data
    })
  }
}

export const initUser = () => {
  const username = window.localStorage.getItem(USER_NAME)
  const token = window.localStorage.getItem(USER_TOKEN)
  if(username !== null && token !== null){
    encryptionService.setToken(token)
    return{
      type: 'LOGIN',
      data: { username, token }
    }
  }else{
    return{
      type: 'LOGOUT'
    }
  }
}

export const logOut = () => {
  window.localStorage.removeItem(USER_NAME)
  window.localStorage.removeItem(USER_TOKEN)
  return {
    type: 'LOGOUT'
  }
}

export default userReducer
