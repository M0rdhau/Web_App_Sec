import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { initUser } from './reducers/userReducer'
import { UserScreen } from './components/UserScreen'
import { Stats } from './components/Stats'
import Notification from './components/presentational/Notification'

const App = () => {
  const data = useSelector(state => state.linkData)
  const dispatch = useDispatch()
  useEffect(() => {
    dispatch(initUser())
  }, [dispatch])

  const user = useSelector(state => state.user)
  const message = useSelector(state => state.notification.text)
  const error = useSelector(state => state.notification.error)
  const notification = { message, error }

  return (
    <div className='mainWrapper'>
      <Notification notification={notification}/>
      <div className="mainBody">
        <UserScreen/>
        {user && data.length > 0 && <Stats/>}
      </div>
    </div>
  )

}

export default App