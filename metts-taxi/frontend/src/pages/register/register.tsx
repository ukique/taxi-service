import {Link, useNavigate} from "react-router-dom"
import "./register.css"
import {useState} from "react";
import axios from "axios";

function Register() {
    const navigate = useNavigate();
    const [error, setError] = useState('')
    const [loading, setLoading] = useState(false);
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const handleCreate = async () => {
        setLoading(true);
        setError("");
        const data = {
            username,
            email,
            password
        }
        try {
            const response = await axios.post(
                "http://localhost:8081/register",
                data
            );
            if (response.status === 201) {
                navigate("/login");
            }
        } catch (err: any) {
            if (err.response?.data?.error) {
                setError(err.response.data.error);
            } else {
                setError("Server Error. Try Later");
            }
        } finally {
            setLoading(false);
        }
    }
    return (
        <div className="auth">
            <div className="card">
                <h1 className="auth-tittle">Register</h1>
                <div className="auth-inputs">
                    <label className="label">USERNAME</label>
                    <input
                        onChange={(e) => setUsername(e.target.value)}
                        className="input" name="username" placeholder="Username"/>
                    <label className="label">EMAIL</label>
                    <input
                        onChange={(e) => setEmail(e.target.value)}
                        className="input" name="email" type="email" placeholder="Email"/>
                    <label className="label">PASSWORD</label>
                    <input
                        onChange={(e) => setPassword(e.target.value)}
                        className="input" name="password" type="password" placeholder="Password"/>
                    <h2 className="auth-error">{error}</h2>
                </div>
                <div className="auth-btns">
                    <button disabled={loading} onClick={handleCreate} className="create-btn">
                        {loading ? "Creating" : "Create"}</button>
                    <button className="google-btn">
                        <img src="https://www.google.com/favicon.ico" alt="google-logo"/>
                        Google
                    </button>
                </div>
                <p className="auth-footer">Already have an account? <Link to="/login">Log In.</Link></p>
            </div>
        </div>
    )
}

export default Register