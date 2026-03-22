import './App.css'
import {Navigate, Route, Routes} from "react-router-dom";
import Orders from "./pages/ordersPage/ordersPage.jsx";
import Drivers from "./pages/driversPage/driversPage.jsx";

function App() {
  return (
    <>
      <Routes>
          <Route path="/" element={<Navigate to="/orders" />} />
          <Route path="/orders" element={<Orders />} />
          <Route path="/drivers" element={<Drivers />} />
      </Routes>
    </>
  )
}

export default App
