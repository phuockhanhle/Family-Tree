import { Button } from "@mui/material";
import { useNavigate } from "react-router-dom";
import { logIn } from "../../components/utils/HelperFuncs";

export const Login = () => {
    const navigate = useNavigate();
    return (
        <>
            <h1> Login Page </h1>
            <Button onClick={() => logIn(navigate)}>Login</Button>
        </>
    )
}
