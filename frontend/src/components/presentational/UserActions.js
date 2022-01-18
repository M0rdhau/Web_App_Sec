import React from "react";
import PropTypes from "prop-types";

export const UserActions = (props) => {
    return (
        <div className="mainParent">
            <div className="encType">
            <h1>Choose your Encryption!</h1>
                {props.encTypes.map((type, index) => (
                    <button name={type.route} key={index} className="encButton" onClick={props.handleEncTypeSwitch}>{type.name}</button>
                ))}
            </div>
            <div className="loginForm">
                <div className="userForm">
                <p>Hello, {props.user.username}</p>
                <button className="logoutButton" onClick={props.handleLogout}>
                    Log out
                </button>
                </div>
                {props.children}
            </div>
        </div>
    );
};

UserActions.propTypes = {
    user: PropTypes.object.isRequired,
    handleLogout: PropTypes.func.isRequired,
    handleEncTypeSwitch: PropTypes.func.isRequired,
    encTypes: PropTypes.array.isRequired,
};
