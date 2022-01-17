import React, { useState } from "react";
import { useForm } from "../hooks/useForm";
import { useDispatch, useSelector } from "react-redux";
import { loginUser, logOut } from "../reducers/userReducer";
import { setNotification } from "../reducers/notificationReducer";
import register from "../services/register";
import Login from "./presentational/Login";
import { UserActions } from "./presentational/UserActions";
import Register from "./presentational/Register";
import { Encryptions } from "./Encryptions";

export const UserScreen = () => {
    const [registering, setRegistering] = useState(false);
    const [loginValues, handleLoginChange, resetLogin] = useForm({
        username: "",
        password: "",
    });
    const [registerValues, handleRegisterChange, resetRegister] = useForm({
        username: "",
        password: "",
        confirmPassword: "",
    });
    const [areEqual, setAreEqual] = useState(true);
    const [encType, setEncType] = useState({
        name: "Caesar",
        route: "/caesar",
    });

    const encTypes = [
        {
            name: "Caesar",
            route: "/caesar",
        },
        {
            name: "Vigenere",
            route: "/vigenere",
        },
        {
            name: "Diffie-Hellmann",
            route: "/diffiehellman",
        },
        {
            name: "RSA - Key Generation",
            route: "/rsa/generate",
        },
        {
            name: "RSA - Encrypt/Decrypt",
            route: "/rsa/generate",
        },
    ];

    const dispatch = useDispatch();

    const user = useSelector((state) => state.user);

    const toggle = () => {
        setRegistering(!registering);
    };

    // login screen code
    const handleLogin = async () => {
        try {
            await dispatch(
                loginUser(loginValues.username, loginValues.password)
            );
        } catch (e) {
            dispatch(setNotification(e.response.headers.error));
        }
        resetLogin();
    };

    //Logged in user panel code
    const handleLogout = () => {
        dispatch(logOut());
    };

    //registration code
    const handleRegister = async () => {
        if (areEqual) {
            try {
                await register(
                    registerValues.username,
                    registerValues.password
                );
                dispatch(setNotification("Registration success!", false));
                resetRegister();
                toggle();
            } catch (e) {
                dispatch(setNotification(e.response.headers.error));
            }
        }
    };

    const handleEncTypeSwitch = (event) => {
        setEncType(
            encTypes.filter((encType) => encType.route === event.target.name)
        );
    };

    const onConfirmPasswordChange = (event) => {
        setAreEqual(event.target.value === registerValues.password);
        handleRegisterChange(event);
    };

    return (
        <>
            {!registering ? (
                user !== null ? (
                    <UserActions
                        user={user}
                        handleLogout={handleLogout}
                        handleEncTypeSwitch={handleEncTypeSwitch}
                        encTypes={encTypes}
                    >
                        <Encryptions encType={encType}/>
                    </UserActions>
                ) : (
                    <Login
                        toggle={toggle}
                        values={loginValues}
                        handleChange={handleLoginChange}
                        handleLogin={handleLogin}
                    />
                )
            ) : (
                <Register
                    toggle={toggle}
                    areEqual={areEqual}
                    values={registerValues}
                    handleChange={handleRegisterChange}
                    handleRegister={handleRegister}
                    onConfirmPasswordChange={onConfirmPasswordChange}
                />
            )}
        </>
    );
};
