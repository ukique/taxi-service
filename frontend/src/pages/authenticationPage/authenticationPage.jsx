import './authenticationPage.css'
import logoWhite from "../../assets/logoWhite.png";
import "../../pages/registerPage/registerPage.css";
import {useNavigate} from "react-router-dom";
import {useState} from "react";

function Authentication() {
    const navigate = useNavigate();
    const [form, setForm] = useState({username: "", email: "", password: ""});
    const [error, setError] = useState(null);
    const handleChange = (e) => {
        setForm({...form, [e.target.name]: e.target.value});
    };
    // against multiple button clicks
    const [isLoading, setIsLoading] = useState(false)

    const handleSubmit = async () => {
        try {
            setIsLoading(true)
            setError(null)
            const response = await fetch("http://localhost:8080/users/authentication", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify(form)
            })
            const result = await response.json()
            if (response.ok) {
                localStorage.setItem("token", result.token)
                localStorage.setItem("username", form.username)
                navigate("/")
            } else {
                setError(result.message)
            }
        } catch (err) {
            setError("Network error. Please try again.")
        } finally {
            setIsLoading(false)
        }
    }

    return (
        <>
            <div className="register-layout">
                <div className="registerBandle">
                    <img className="logoWhite" src={logoWhite} alt="logoWhite"/>
                    <div className="registerBandleText">
                        <h1>TAXI</h1>
                        <h1 id="h1-outline">CONTROL</h1>
                        <h1>PANEL</h1>
                    </div>
                    <div className="registerBandleInfo">
                        <p>
                            Real-time driver tracking and ride<br/>
                            management for modern dispatch <br/>
                            operations.
                        </p>
                    </div>
                </div>
                <div className="register">
                    <div className="register-text">
                        <h1>Authentication</h1>
                        <h3>Enter your data to login</h3>
                    </div>
                    <p className="register-input-text">USERNAME</p>
                    <input name="username" value={form.username} onChange={handleChange} placeholder="Username"/>
                    <p className="register-input-text">EMAIL</p>
                    <input name="email" type="email" value={form.email} onChange={handleChange} placeholder="Email"/>
                    <p className="register-input-text">PASSWORD</p>
                    <input name="password" type="password" value={form.password} onChange={handleChange}
                           placeholder="Password"/>
                    {error && <p className="register-error">{error}</p>}
                    <button onClick={handleSubmit} disabled={isLoading} className="register-button">
                        {isLoading ? "Creating..." : "Sign In"}
                    </button>
                    <p className="register-auth-request">Don't have an account? <a onClick={() => navigate("/users/register")}>Sign up</a></p>
                </div>
            </div>
            </>
    )
}

export default Authentication