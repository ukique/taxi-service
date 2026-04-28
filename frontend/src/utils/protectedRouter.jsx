import { Outlet, Navigate } from "react-router-dom"
import { useEffect, useState } from "react"
import { refreshAccessToken } from "../api/authApi.js"

const ProtectedRoutes = () => {
    const [isAuthenticated, setIsAuthenticated] = useState(null)

    useEffect(() => {
        refreshAccessToken()
            .then(() => setIsAuthenticated(true))
            .catch(() => setIsAuthenticated(false))
    }, [])

    if (isAuthenticated === null) return <div>Loading...</div>

    return isAuthenticated ? <Outlet/> : <Navigate to="/users/authentication"/>
}

export default ProtectedRoutes