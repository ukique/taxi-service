import "./driversPage.css"
import Header from "../../components/header/header.jsx";
import RegisterDriverButton from "../../components/drivers/registerDriver.jsx";

function Drivers() {
    return (
        <>
            <div  className="drivers-page">
            <Header/>
            <div className="drivers-page-main">
                <h1>Drivers</h1>
                <RegisterDriverButton/>
            </div>
            <h4>Manage driver profiles and monitor status</h4>
            </div>
        </>
    )
}

export default Drivers