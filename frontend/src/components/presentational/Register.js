import React from 'react'
import PropTypes from 'prop-types'

const Register =
({ toggle, areEqual, values, handleChange, onConfirmPasswordChange, handleRegister }) => {

  return(
    <div className="loginForm">
      {!areEqual && <p> Passwords must match! </p>}
      <label htmlFor='username'>Username: </label>
      <input className='userNameInput' name='username' value={values.username} onChange={handleChange} />
      <label htmlFor='password'>Password: </label>
      <input
        className='passwordInput'
        name='password'
        type='password'
        value={values.password}
        onChange={handleChange} />
      <label htmlFor='confirmPassword'>Confirm Password:</label>
      <input
        className='confirmPassword'
        name='confirmPassword'
        type='password'
        value={values.confirmPassword}
        onChange={onConfirmPasswordChange} />
      <button className='registerButton' onClick={handleRegister}>Register</button>
      <button className='toggleButton' onClick={toggle}>Or Log In</button>
    </div>
  )

}

Register.propTypes = {
  toggle: PropTypes.func.isRequired,
  areEqual: PropTypes.bool.isRequired,
  values: PropTypes.object.isRequired,
  handleChange: PropTypes.func.isRequired,
  onConfirmPasswordChange: PropTypes.func.isRequired,
  handleRegister: PropTypes.func.isRequired
}

export default Register
