import axios from "axios";
import Cookies from "js-cookie";
import { NavigateFunction } from "react-router-dom";
import { back_end_url, cookie_login_name, cookie_login_value, isLocal } from "../../app/constants";

type cookieProps = {
  nom: string,
  path: string,
  domain: string,
  value: string,
}

export function goToExternalPage(link: string | undefined) {
  if (link)
    window.location.assign(link);
};

export function openNewTabWithLink(link: string | undefined, cookies?: cookieProps[]) {
  if (link) {
    if (cookies) {
      cookies.forEach((cookie: cookieProps) => {
        Cookies.set(cookie.nom, cookie.value, { path: cookie.path, domain: cookie.domain });
        // console.log(cookie.nom);
      })
    }
    window.open(link, "_blank");
  }
  else window.open();
}

export function isObjectEmpty(object: any) {
  return (Object.keys(object).length === 0);
}

export function isArrayEmpty(arr: Array<any>) {
  return (arr.length === 0);
}

// In Local, check if login cookie exists, else navigate to loginPage ('/')
export function cookieCheck(navigate: any) {
  if (isLocal && !Cookies.get(cookie_login_name)) navigate('/');
}

// In Local, check if login cookie exists, else logout
export function securityCheck(navigate: any) {
  if (isLocal && !Cookies.get(cookie_login_name)) logOut(navigate);
}

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