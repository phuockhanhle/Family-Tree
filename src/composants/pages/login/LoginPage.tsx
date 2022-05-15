import axios from "axios";
import Cookies from "js-cookie";
import { useEffect } from "react";
import { isIE } from "react-device-detect";
import { Navigate, useNavigate } from "react-router-dom";
import { back_end_url, cookie_login_name, isLocal } from "../../../app/constants";
import { cookieCheck, logOut } from "../../utils/helperFunc";
import { Login } from "./Login";

export const LoginPage = () => {
    // useEffect(() => {
    // }, [])

    return <Login />

}