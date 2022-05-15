import { Button } from "@material-ui/core";
import { useEffect } from "react";
import { isIE } from "react-device-detect";
import { useLocation, useNavigate } from "react-router-dom";
import { isObjectEmpty, logOut, securityCheck } from "../../utils/helperFunc";
import { DashBoard } from "./DashBoard";

/** DashBoardPage ('/dashboard'): main page*/
export const DashBoardPage = () => {
    const navigate = useNavigate();
    const parent = {
        dad: {
            name: 'Adam',
            relation: ''
        },
        mom: {
            name: 'Anna',
            relation: ''
        }
    }
    const child = {
        name: 'Alex',
        gender: 'Male'
    };

    if (parent)
        return (
            <div>
                <DashBoard parent={parent} child={child}
                />
                <Button onClick={() => logOut(navigate)}>logout</Button>
            </div>
        )
}