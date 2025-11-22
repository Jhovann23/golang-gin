// import FC from react
import {type FC, useState, useContext,type FormEvent } from 'react';

//import hook useNavigate from react router
import { useNavigate } from "react-router";

//import custom  hook useLogin from hooks
import { useLogin } from "../../hooks/auth/useLogin";

//import js-cookie
import Cookies from 'js-cookie'

//import context
import { AuthContext } from '../../context/AuthContext';

//interface for validation errors
interface ValidationErrors {
    [key: string]: string;
}

const Login: FC = () => {

    const navigate = useNavigate();

    //mutate login
    const {mutate, isPending} = useLogin();

    //authenticated ?
    const {setIsAuthenticated} = useContext(AuthContext)!;

    //binding 
    const [username, setUsername] = useState<string>('');
    const [password, setPassword] = useState<string>('');

    //errors
    const [errors, setErrors] = useState<ValidationErrors>({})

    const handleLogin =  async (e: FormEvent) => {
        e.preventDefault();

        mutate({
            username,
            password
        }, {
            onSuccess: (data: any) => {
                //set token cookie
                Cookies.set('token', data.data.token);

                //set token user
                Cookies.set('user', JSON.stringify({
                    id: data.data.id,
                    name: data.data.name,
                    username: data.data.name,
                    email: data.data.email
                }))

                setIsAuthenticated(true);

                //redirect ke dashboard
                navigate('/admin/dashboard');
            },
            onError: (error: any) => {
                setErrors(error.response.data.errors)
            } 
        })
    } 


    return (
        <div className="row justify-content-center mt-5">
            <div className="col-md-4">
                <div className="card border-0 rounded-4 shadow-sm">
                    <div className="card-body">
                        <h4 className='fw-bold text-center'>LOGIN</h4>
                        <hr />
                        {errors.Error && <div className="alert alert-danger mt-2 rounded-4">Username or Password is incorrect</div>}
                        <form onSubmit={handleLogin}>
                            <div className="form-group mb-3">
                                <label className="mb-1 fw-bold">Username</label>
                                <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} className="form-control" placeholder="Username" />
                                {errors.Username && <div className="alert alert-danger mt-2 rounded-4">{errors.Username}</div>}
                            </div>

                            <div className="form-group mb-3">
                                <label className="mb-1 fw-bold">Password</label>
                                <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} className="form-control"
                                    placeholder="Password" />
                                {errors.Password && <div className="alert alert-danger mt-2 rounded-4">{errors.Password}</div>}
                            </div>
                            <button type="submit" className="btn btn-primary w-100 rounded-4" disabled={isPending}>
                                {isPending ? 'Loading...' : 'LOGIN'}
                            </button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Login;