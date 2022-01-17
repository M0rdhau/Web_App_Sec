import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { initUser } from './reducers/userReducer'
import { UserScreen } from './components/UserScreen'
import Notification from './components/presentational/Notification'

const App = () => {
  const dispatch = useDispatch()
  useEffect(() => {
    dispatch(initUser())
  }, [dispatch])

  const message = useSelector(state => state.notification.text)
  const error = useSelector(state => state.notification.error)
  const notification = { message, error }

  return (
    <div className='mainWrapper'>
      <Notification notification={notification}/>
      <div className="mainBody">
        <UserScreen/>
      </div>
    </div>
  )

}

export default App