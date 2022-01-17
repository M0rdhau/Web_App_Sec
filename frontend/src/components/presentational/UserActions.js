import React from 'react'
import PropTypes from 'prop-types'

export const UserActions =
(props) => {
  return(
    <div className="loginForm">
      <p>Hello, {props.user.username}</p>
      <button className='logoutButton' onClick={props.handleLogout} >Log out</button>
      {props.children}
    </div>
  )
}

UserActions.propTypes = {
  user: PropTypes.object.isRequired,
  handleLogout: PropTypes.func.isRequired,
}