import './App.css'
import {Navigate, Route, Routes} from "react-router-dom";
import Orders from "./pages/ordersPage/ordersPage.jsx";
import Drivers from "./pages/driversPage/driversPage.jsx";
import Registration from "./pages/registerPage/registerPage.jsx";
import Authentication from "./pages/authenticationPage/authenticationPage.jsx";
import ProtectedRouter from "./utils/protectedRouter.jsx";
import {useEffect} from "react";
import ws from "./services/websocket.js";
import OrderDetails from "./components/orders/order-details/orderDetails.jsx";

function App() {
    useEffect(() => {
        ws.connect("ws://localhost:8080/ws");
    }, []);
    return (
        <Routes>
            <Route path="/users/register" element={<Registration/>}/>
            <Route path="/users/authentication" element={<Authentication/>}/>
            <Route element={<ProtectedRouter/>}>
                <Route path="/" element={<Navigate to="/orders"/>}/>
                <Route path="/orders/page/:id" element={<OrderDetails/>}/>
                <Route path="/orders" element={<Orders/>}/>
                <Route path="/drivers" element={<Drivers/>}/>
            </Route>
        </Routes>
    )
}

export default App