import "./createOrder.css"
import {useState} from "react";

function CreateOrderButton() {
    const [openPopup, setOpenPopup] = useState(false)
    const [form, setForm] = useState({user_id: 0})
    const [error, setError] = useState(null);
    const handleChange = (e) => {
        setForm({...form, [e.target.name]: parseInt(e.target.value, 10) || 0})
    }

    const [isLoading, setIsLoading] = useState(false)
    const handleSubmit = async () => {
        try {
            setIsLoading(true)
            setError(null)
            const response = await fetch("http://localhost:8080/orders", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify(form)
            })
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
        <>
            <div className="create-driver-button">
                <button onClick={() => setOpenPopup(true)}>Create Order</button>
            </div>
            {openPopup && <div className="create-driver-menu">
                <h1>Create Order</h1>
                <h2>Enter User ID</h2>
                <div className="orders-input">
                    <input type="number" name="user_id" onChange={handleChange} disabled={isLoading}
                           placeholder="UserID"/>
                </div>
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