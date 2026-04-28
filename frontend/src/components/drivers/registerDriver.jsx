import "./registerDriver.css"
import {useState} from "react";
import {refreshAccessToken} from "../../api/authApi.js";

function RegisterDriverButton() {
    const [openPopup, setOpenPopup] = useState(false)
    const [form, setForm] = useState({username: ""})
    const [error, setError] = useState(null);
    const handleChange = (e) => {
        setForm({...form, [e.target.name]: e.target.value})
    }
    const [isLoading, setIsLoading] = useState(false)
    const handleSubmit = async () => {
        try {
            setIsLoading(true)
            setError(null)
            let response = await fetch("http://localhost:8080/drivers/create", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify(form),
                credentials: "include"
            })
            if (response.status === 401) {
                await refreshAccessToken()
                response = await fetch("http://localhost:8080/drivers/create",{
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify(form),
                    credentials: "include"
                })
            }
            const result = await response.json()
            if (result.ok) {
                setOpenPopup(false)
            } else {
                setError(result.message)
            }
        } catch (err) {
            setError("Network error. Please try again.")
        } finally {
            setIsLoading(false)
        }
    }

    return (
        <div className="add-driver">
            <button className="add-driver-button" onClick={() => setOpenPopup(true)}>+ Add
                Driver
            </button>
            {
                openPopup &&
                <div className="driver-register">
                    <h2>Add Driver</h2>
                    <p>Driver Name</p>
                    <input name="username" onChange={handleChange} disabled={isLoading} placeholder="Driver Name"/>
                    {error && <p className="register-error">{error}</p>}
                    <div className="drivers-register-buttons">
                        <button className="add-driver-cancel"
                                onClick={() => setOpenPopup(false)}>Cancel
                        </button>
                        <button className="add-driver-save" onClick={handleSubmit} disabled={isLoading}>
                            {isLoading ? "Creating..." : "Save Driver"}
                        </button>
                    </div>
                </div>
            }
        </div>
    )
}

export default RegisterDriverButton