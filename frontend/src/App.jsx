import './App.css'
import {Navigate, Route, Routes} from "react-router-dom";
import Orders from "./pages/ordersPage/ordersPage.jsx";
import Drivers from "./pages/driversPage/driversPage.jsx";
import Registration from "./pages/registerPage/registerPage.jsx";
import Authentication from "./pages/authenticationPage/authenticationPage.jsx";
import ProtectedRouter from "./utils/protectedRouter.jsx";
import {useEffect} from "react";
import ws from "./services/websocket.js";
import OrderDetailsLocation from "./pages/ordersPage/orderDetails/order-details-location.jsx";
import OrderDetailsInfo from "./pages/ordersPage/orderDetails/order-details-info.jsx";
import DriversHistory from "./pages/driversPage/drivers-history/drivers-history.jsx";

function App() {
    useEffect(() => {
        ws.connect("ws://localhost:8080/ws");
    }, []);
    return (
        <Routes>
            <Route path="/users/register" element={<Registration/>}/>
            <Route path="/users/authentication" element={<Authentication/>}/>
            <Route element={<ProtectedRouter/>}>
                <Route path="/" element={<Navigate to="/orders/1"/>}/>
                <Route path="/orders/page/:id" element={<OrderDetailsInfo/>}/>
                <Route path="/orders/page/:id/location" element={<OrderDetailsLocation/>}/>
                <Route path="/orders/:id" element={<Orders/>}/>
                <Route path="/drivers" element={<Drivers/>}/>
                <Route path="/drivers/:id/page/:pageID" element={<DriversHistory/>}/>
            </Route>
        </Routes>
    )
}

export default App