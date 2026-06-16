import {Link} from "react-router-dom";
import "../register/register.css"

function LogIn() {
    return (
        <div className="auth">
            <div className="card">
                <h1 className="auth-tittle">Login</h1>
                <div className="auth-inputs">
                    <label className="label">USERNAME</label>
                    <input className="input" name="username" placeholder="Username"/>
                    <label className="label">EMAIL</label>
                    <input className="input" name="email" type="email" placeholder="Email"/>
                    <label className="label">PASSWORD</label>
                    <input className="input" name="password" type="password" placeholder="Password"/>
                </div>
                <div className="auth-btns">
                    <button className="create-btn">Log In</button>
                    <button className="google-btn">
                        <img src="https://www.google.com/favicon.ico" alt="google-logo"/>
                        Google
                    </button>
                </div>
                <p className="auth-footer">Don't have an account? <Link to="/register">Register.</Link></p>
            </div>
        </div>
    )
}

export default LogIn