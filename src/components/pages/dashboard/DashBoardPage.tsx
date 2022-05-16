import Button from "@mui/material/Button";
import { useEffect, useState } from "react";
import { isIE } from "react-device-detect";
import { useLocation, useNavigate } from "react-router-dom";
import { SimpleDialog } from "../../boxes/dialogBox/DialogBox";
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
            uid: 'adfsgfergcfzerfdsfzeghgkl',
            show: true
        },
        partner: {
            name: 'Olivia',
            gender: 'Female',
            age: 25,
            uid: 'adfsgfergcfzesdzfsfeeghgkl',
            show: true
        },
        parents: {
            total: 2,
            details: [
                {
                    name: 'Adam',
                    gender: 'Male',
                    age: 40,
                    uid: 'adfsgfergcfzerfdsfzeghgkl',
                    show: true
                },
                {
                    name: 'Eva',
                    gender: 'Female',
                    age: 37,
                    uid: 'adfsgfergcfzerfdsfzeghgkl',
                    show: true
                }]
        },
        children: {
            total: 2,
            details: [
                {
                    name: 'Alex',
                    gender: 'Male',
                    age: 5,
                    uid: 'adfsgfergcfzerfdsfzeghgkl',
                    show: true
                },
                {
                    name: 'Anna',
                    gender: 'Female',
                    age: 4,
                    uid: 'adfsgfergcfzerfdsfzeghgkl',
                    show: true
                }
            ]
        },
        siblings: {
            total: 2,
            details: [
                {
                    name: 'Bob',
                    gender: 'Male',
                    age: 28,
                    uid: 'adfsgfergcfzerfdsfzeghgkl',
                    show: true
                },
                {
                    name: 'Elena',
                    gender: 'Female',
                    age: 20,
                    uid: 'adfsgfergcfzerfdsfzeghgkl',
                    show: true
                }
            ]
        }
    }

    const options = [
        'Add person',
        'Remove person'
    ]

    const [open, setOpen] = useState(false);
    const [selectedValue, setSelectedValue] = useState(options[0]);

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = (value: string) => {
        setOpen(false);
        setSelectedValue(value);
    };

    return (
        <div>
            <DashBoard person={person} />
            <br />
            <br />
            <Button variant="contained" onClick={handleClickOpen}>+</Button>
            <SimpleDialog
                selectedValue={selectedValue}
                open={open}
                onClose={handleClose}
            />
            <br />
            <br />
            <Button variant="contained" onClick={() => logOut(navigate)}>logout</Button>
        </div >
    )
}