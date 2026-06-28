import {Link, useNavigate} from "react-router-dom";
import "../register/register.css"
import axios from "axios"
import {useState} from "react";

function LogIn() {
    const navigate = useNavigate()
    const [error, setError] = useState("")
    const [username, setUsername] = useState("")
    const [loading, setLoading] = useState(false);
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const handleCreate = async () => {
        setLoading(true);
        const data = {
            username,
            email,
            password,
        };
        try {
            const response = await axios.post(
                "http://localhost:8081/login",
                data
            );
            if (response.status == 200) {
                localStorage.setItem("username", response.data)
                navigate("/home")
            }
        } catch (err: any) {
            if (err.response?.data?.error) {
                setError(err.response.data.error)
            } else {
                setError("Server Error. Try later.")
            }
        } finally {
            setLoading(false);
        }

    }
    return (
        <div className="auth">
            <div className="card">
                <h1 className="auth-tittle">Login</h1>
                <div className="auth-inputs">
                    <label className="label">USERNAME</label>
                    <input className="input" onChange={(e) => setUsername(e.target.value)} name="username"
                           placeholder="Username"/>
                    <label className="label">EMAIL</label>
                    <input className="input" onChange={(e) => setEmail(e.target.value)} name="email" type="email"
                           placeholder="Email"/>
                    <label className="label">PASSWORD</label>
                    <input className="input" onChange={(e) => setPassword(e.target.value)} name="password"
                           type="password" placeholder="Password"/>
                    <h2 className="auth-error">{error}</h2>
                </div>
                <div className="auth-btns">
                    <button disabled={loading} onClick={handleCreate} className="create-btn">
                        {loading ? "Loading..." : "Log In"}</button>
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