import Button from "@mui/material/Button";
import { useEffect } from "react";
import { isIE } from "react-device-detect";
import { useLocation, useNavigate } from "react-router-dom";
import { isObjectEmpty, logOut, securityCheck } from "../../utils/helperFunc";
import { DashBoard } from "./DashBoard";

/** DashBoardPage ('/dashboard'): main page*/
export const DashBoardPage = () => {
    const navigate = useNavigate();
    const person = {
        id: {
            name: 'Sirius',
            gender: 'Male',
            age: 27,
            uid: 'adfsgfergcfzerfdsfzeghgkl'
        },
        partner: {
            name: 'Olivia',
            gender: 'Female',
            age: 25
        },
        parents: {
            total: 2,
            details: [
                {
                    name: 'Adam',
                    gender: 'Male',
                    age: 40
                },
                {
                    name: 'Eva',
                    gender: 'Female',
                    age: 37
                }]
        },
        children: {
            total: 2,
            details: [
                {
                    name: 'Alex',
                    gender: 'Male',
                    age: 5
                },
                {
                    name: 'Anna',
                    gender: 'Female',
                    age: 4
                }
            ]
        },
        siblings: {
            total: 2,
            details: [
                {
                    name: 'Bob',
                    gender: 'Male',
                    age: 28
                },
                {
                    name: 'Elena',
                    gender: 'Female',
                    age: 20
                }
            ]
        }
    }

    return (
        <div>
            <DashBoard person={person} />
            <Button variant="contained" onClick={() => logOut(navigate)}>logout</Button>
        </div >
    )
}