import './App.css'
import {Route, Routes} from "react-router-dom";
import Register from "./pages/register/register.tsx";
import Home from "./pages/home/home.tsx";
import Login from "./pages/login/login.tsx";

function App() {
    return (
        <div>
            <Routes>
                <Route path={"/"} element={<Home/>}/>
                <Route path={"/register"} element={<Register/>}/>
                <Route path={"/login"} element={<Login/>}/>
            </Routes>
        </div>
    )
}

export default App
