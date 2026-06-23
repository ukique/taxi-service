import Header from "../../components/header/header.tsx";
import './main.css'
import {useNavigate} from "react-router-dom";

function MainPage() {
    const navigate = useNavigate();
    return (
        <>
            <Header/>
            <div className="home">
                <video className="home-intro" autoPlay muted loop playsInline>
                    <source src="/videos/metts-main-porshe.mp4" type="video/mp4"/>
                </video>
                <div className="home-content">
                    <h1 className="home-tittle">Go anywhere with <br/>metts</h1>
                    <p className="home-description">metts is a global premium taxi service <br/> designed for those who
                        value comfort, reliability, and style</p>
                    <button className="home-button" onClick={() => navigate("/register")}>Start Now</button>
                </div>
            </div>
        </>
    )
}

export default MainPage