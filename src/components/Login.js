// https://blog.logrocket.com/handling-user-authentication-redux-toolkit/

import { useForm } from 'react-hook-form'
import { useState, useContext, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Input from "./form/Input";

// Redux
import { useDispatch, useSelector } from 'react-redux'
import { userLogin } from '../redux/actions/userAction'


const Login = () => {
    // Redux
    const { loading, userInfo, error } = useSelector((state) => state.user)
    const dispatch = useDispatch()
    const { register, handleSubmit } = useForm();
  
    
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    
    const [ jwtToken, setJwtToken ] = useState("");
    const [alertMessage, setAlertMessage] = useState("");
    const [alertClassName, setAlertClassName] = useState("d-none");
    // const { setAlertClassName } = useOutletContext();
    // const { setAlertMessage } = useOutletContext();
    // const { toggleRefresh } = useOutletContext();
    
    const navigate = useNavigate();
    // redirect authenticated user to profile screen
    useEffect(() => {
        if (userInfo) {
            navigate('/keywords')
        }
    }, [navigate, userInfo])
    
    // const handleSubmit = (event) => {
    //     event.preventDefault();
        
    //     // build the request payload
    //     let payload = {
    //         email: email,
    //         password: password,
    //     }
        
    //     const requestOptions = {
    //         method: "POST",
    //         headers: {
    //             'Content-Type': 'application/json'
    //         },
    //         credentials: 'include',
    //         body: JSON.stringify(payload),
    //     }
        
    //     fetch(`/login`, requestOptions)
    //     .then((response) => response.json())
    //     .then((data) => {
    //         if (data.error) {
    //             setAlertClassName("alert-danger");
    //             setAlertMessage(data.message);
    //         } else {
    //             if(data.token) {
    //                 localStorage.setItem("user", JSON.stringify(data.token))
    //             }
    //             setJwtToken(data.token);
    //             // setAlertClassName("d-none");
    //             // setAlertMessage("");
    //             // toggleRefresh(true);
    //             navigate("/keywords");
    //         }
    //     })
    //     .catch(error => {
    //         // setAlertClassName("alert-danger");
    //         // setAlertMessage(error);
    //     })
        
    // }
    const submitForm = (data) => {
      console.log(data)
      dispatch(userLogin(data))
    }

    return(
        <div className="col-md-6 offset-md-3">
            <h2>Login</h2>
            <hr />

            <form onSubmit={handleSubmit(submitForm)}>
                <div className='form-group'>
                    <label htmlFor='email'>Email</label>
                    <input
                    type='email'
                    className='form-input form-control'
                    {...register('email')}
                    required
                    />
                </div>
                <div className='form-group'>
                    <label htmlFor='password'>Password</label>
                    <input
                    type='password'
                    className='form-input form-control'
                    {...register('password')}
                    required
                    />
                </div>
                {/* <Input
                    title="Email Address"
                    type="email"
                    className="form-control"
                    name="email"
                    autoComplete="email-new"
                    {...register('email')}
                    onChange={(event) => setEmail(event.target.value)}
                />

                <Input
                    title="Password"
                    type="password"
                    className="form-control"
                    name="password"
                    autoComplete="password-new"
                    {...register('password')}
                    onChange={(event) => setPassword(event.target.value)}
                /> */}

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