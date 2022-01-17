import React from 'react'
import PropTypes from 'prop-types'

const Notification = ({ notification }) => {
  if(notification.message === ''){
    return null
  }
  return (
    <div className={notification.error ? 'error' : 'notification'}>
      <p>{notification.message}</p>
    </div>
  )
}

Notification.propTypes = {
  notification: PropTypes.object.isRequired
}


export default Notification
