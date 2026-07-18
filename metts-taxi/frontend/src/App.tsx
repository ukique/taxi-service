import './App.css'
import {Route, Routes} from "react-router-dom";
import Register from "./pages/register/register.tsx";
import Login from "./pages/login/login.tsx";
import MainPage from "./pages/main/main.tsx";
import Home from "./pages/home/home.tsx";

function App() {
    return (
        <div>
            <Routes>
                <Route path={"/"} element={<MainPage/>}/>
                <Route path={"/register"} element={<Register/>}/>
                <Route path={"/login"} element={<Login/>}/>
                <Route path={"/home"} element={<Home/>}/>
            </Routes>
        </div>
    )
}

export default App
