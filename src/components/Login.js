import { useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { userToken } from "../App";
import Input from "./form/Input";

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    
    const [ jwtToken, setJwtToken ] = useState("");
    const [alertMessage, setAlertMessage] = useState("");
    const [alertClassName, setAlertClassName] = useState("d-none");
    // const { setAlertClassName } = useOutletContext();
    // const { setAlertMessage } = useOutletContext();
    // const { toggleRefresh } = useOutletContext();

    const navigate = useNavigate();

    const handleSubmit = (event) => {
        event.preventDefault();
        
        // build the request payload
        let payload = {
            email: email,
            password: password,
        }

        const requestOptions = {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            credentials: 'include',
            body: JSON.stringify(payload),
        }

        fetch(`/login`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    setAlertClassName("alert-danger");
                    setAlertMessage(data.message);
                } else {
                    if(data.token) {
                        localStorage.setItem("user", JSON.stringify(data.token))
                    }
                    setJwtToken(data.token);
                    // setAlertClassName("d-none");
                    // setAlertMessage("");
                    // toggleRefresh(true);
                    navigate("/keywords");
                }
            })
            .catch(error => {
                // setAlertClassName("alert-danger");
                // setAlertMessage(error);
            })
    }

    return(
        <div className="col-md-6 offset-md-3">
            <h2>Login</h2>
            <hr />

            <form onSubmit={handleSubmit}>
                <Input
                    title="Email Address"
                    type="email"
                    className="form-control"
                    name="email"
                    autoComplete="email-new"
                    onChange={(event) => setEmail(event.target.value)}
                />

                <Input
                    title="Password"
                    type="password"
                    className="form-control"
                    name="password"
                    autoComplete="password-new"
                    onChange={(event) => setPassword(event.target.value)}
                />

                <hr />

                <input 
                    type="submit"
                    className="btn btn-primary"
                    value="Login"
                />


            </form>
        </div>
    )
}

export default Login;