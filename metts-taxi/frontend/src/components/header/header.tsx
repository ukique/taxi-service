import "./header.css"
import {useNavigate} from "react-router-dom";

function Header() {
    const navigate = useNavigate();

    return (
        <>
            <header className="header">
                <div className="header-logo" onClick={() => navigate("/")}>
                    <h1>metts.</h1>
                </div>
                <button className="header-reg-btn" onClick={() => navigate("/register")}>register</button>
            </header>
        </>
    )
}

export default Header