import './App.css'
import {Route, Routes} from "react-router-dom";
import Register from "./pages/register/register.tsx";
import Login from "./pages/login/login.tsx";
import MainPage from "./pages/main/main.tsx";

function App() {
    return (
        <div>
            <Routes>
                <Route path={"/"} element={<MainPage/>}/>
                <Route path={"/register"} element={<Register/>}/>
                <Route path={"/login"} element={<Login/>}/>
            </Routes>
        </div>
    )
}

export default App
