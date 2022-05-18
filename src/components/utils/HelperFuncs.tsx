import axios from "axios";
import Cookies from "js-cookie";
import { NavigateFunction } from "react-router-dom";
import { back_end_url, cookie_login_name, cookie_login_value, isLocal } from "../../app/constants";


export function logOut(navigate: NavigateFunction) {
    Cookies.remove(cookie_login_name);
    navigate('/');
  }

export function logIn(navigate: NavigateFunction) {
  Cookies.set(cookie_login_name, cookie_login_value);
    navigate('/dashboard', {
      state: {
        parent: '',
        child: '',
        gender: ''
      }
    })
}