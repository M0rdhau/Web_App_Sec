import React from "react";
import PropTypes from "prop-types";

export const UserActions = (props) => {
    return (
        <>
            <div className="encType">
                {props.encTypes.map((type, index) => (
                    <button name={type.route} key={index} className="encButon" onClick={props.handleEncTypeSwitch}>{type.name}</button>
                ))}
            </div>
            <div className="loginForm">
                <p>Hello, {props.user.username}</p>
                <button className="logoutButton" onClick={props.handleLogout}>
                    Log out
                </button>
                {props.children}
            </div>
        </>
    );
};

UserActions.propTypes = {
    user: PropTypes.object.isRequired,
    handleLogout: PropTypes.func.isRequired,
    handleEncTypeSwitch: PropTypes.func.isRequired,
    encTypes: PropTypes.array.isRequired,
};
