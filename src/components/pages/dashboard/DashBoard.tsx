import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { NameBox } from "../../boxes/nameBox/NameBox";
import Typography from "@mui/material/Typography";
import { Grid, Box, Tooltip, Button } from "@mui/material";
import { FamilyBox } from "../../boxes/familyBox/FamilyBox";


type dashBoardProps = {
    person: any
}

export const DashBoard = (props: dashBoardProps) => {

    return (
        <FamilyBox person={props.person}/>
    );
}
