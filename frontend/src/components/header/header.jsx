import './header.css'
import logo from "../../assets/logo.png"
import {useLocation, useNavigate} from "react-router-dom";

function Header() {
    const location = useLocation()
    const navigate = useNavigate()
    return (
        <>
            <header className="header">
                <a onClick={() => navigate("/orders")}><img className="logo" src={logo} alt="logo"/></a>
                <nav>
                    <button className={`headerButton ${location.pathname === "/orders" ? "active" : ""}`} onClick={() => navigate("/orders")}>Orders
                    </button>
                    <button className={`headerButton ${location.pathname === "/drivers" ? "active" : ""}`} onClick={() => navigate("/drivers")}>Drivers
                    </button>
                </nav>
            </header>
        </>
    )
}

export default Header
