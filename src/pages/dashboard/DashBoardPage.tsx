import Button from "@mui/material/Button";
import { useEffect, useState } from "react";
import { isIE } from "react-device-detect";
import { useLocation, useNavigate } from "react-router-dom";

import { DashBoard } from "./DashBoard";

/** DashBoardPage ('/dashboard'): main page*/
export const DashBoardPage = () => {
    
    return (
        <DashBoard/>    
    )
}