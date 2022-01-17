const TIME = 5000

const notificationReducer = (state = { id: 0, text: '' }, action) => {
  switch(action.type) {
    case 'NEW_NOTIFICATION':
      if(state.id !== 0) window.clearTimeout(state.id)
      return {
        id: action.id,
        text: action.text,
        error: action.error
      }
    case 'CLEAR':
      return {
        id: 0,
        text: '',
        error: false
      }
    default:
      return state
  }
}

export const updateNotif = (text, id, error) => {
  return {
    type: 'NEW_NOTIFICATION',
    id,
    text,
    error
  }
}

export const clearNotif = () => {
  return {
    type: 'CLEAR'
  }
}

export const setNotification = (text, error) => {
  return async dispatch => {
    const timeID = window.setTimeout(() => {
      dispatch(
        clearNotif()
      )
    }, TIME)
    dispatch(updateNotif(text, timeID, error))
  }
}

export default notificationReducer
