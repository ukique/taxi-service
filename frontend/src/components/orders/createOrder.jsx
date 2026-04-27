import "./createOrder.css"
import {useState} from "react";
import {refreshAccessToken} from "../../api/authApi.js";

function CreateOrderButton() {
    const [openPopup, setOpenPopup] = useState(false)
    const [error, setError] = useState(null);

    const [isLoading, setIsLoading] = useState(false)
    const handleSubmit = async () => {
        try {
            setIsLoading(true)
            setError(null)
            let response = await fetch("http://localhost:8080/orders", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                credentials: "include"
            })

            if (response.status === 401) {
                await refreshAccessToken()
                response = await fetch("http://localhost:8080/orders", {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
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
            setError("Something went wrong, please try again.")
        } finally {
            setIsLoading(false)
        }
    }
    return (
        <>
            <div className="create-driver-button">
                <button onClick={() => setOpenPopup(true)}>Create Order</button>
            </div>
            {openPopup && <div className="create-driver-menu">
                <h1>Create Order</h1>
                <h2>Do you want to create order?</h2>
                {error && <p className="register-error">{error}</p>}
                <div className="create-order-menu-buttons">
                    <button id="create-order-cancel" onClick={() => setOpenPopup(false)}>Cancel</button>
                    <button id="create-order-button" onClick={handleSubmit} disabled={isLoading}>
                        {isLoading ? "Creating..." : "Create Order"}
                    </button>
                </div>
            </div>
            }
        </>
    )
}

export default CreateOrderButton