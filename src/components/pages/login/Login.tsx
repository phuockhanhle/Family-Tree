import { useNavigate } from "react-router-dom";
import { logIn } from "../../utils/HelperFuncs";

export const Login = () => {
    const navigate = useNavigate();
    return (
        <div>
            <h1> Login Page </h1>
            <button onClick={() => logIn(navigate)}>Login</button>
        </div>
    )
}
