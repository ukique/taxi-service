import './App.css'
import {Navigate, Route, Routes} from "react-router-dom";
import Orders from "./pages/ordersPage/ordersPage.jsx";
import Drivers from "./pages/driversPage/driversPage.jsx";
import Registration from "./pages/registerPage/registerPage.jsx";
import Authentication from "./pages/authenticationPage/authenticationPage.jsx";

function App() {
    return (
        <>
            <Routes>
                <Route path="/" element={<Navigate to="/orders"/>}/>
                <Route path="/orders" element={<Orders/>}/>
                <Route path="/drivers" element={<Drivers/>}/>
                <Route path="/users/register" element={<Registration/>}/>
                <Route path="/users/authentication" element={<Authentication/>}/>
            </Routes>
        </>
    )
}

export default App
