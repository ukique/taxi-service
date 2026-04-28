import {useEffect, useState} from "react"
import "./allDrivers.css"
import {refreshAccessToken} from "../../api/authApi.js";

function AllDriversTable() {
    const [drivers, setDrivers] = useState([])
    const [error, setError] = useState(null)
    const [filterText, setFilterText] = useState("")
    const [isLoading, setIsLoading] = useState(false)

    // popups
    const [openPopupChangeStatus, setOpenPopupChangeStatus] = useState(false)
    const [openPopupChangeName, setOpenPopupChangeName] = useState(false)
    const [openPopupDeleteDriver, setOpenPopupDeleteDriver] = useState(false)

    // forms
    const [statusForm, setStatusForm] = useState({id: "", status: "driving"})
    const [nameForm, setNameForm] = useState({id: "", username: ""})
    const [deleteForm, setDeleteForm] = useState({id: ""})

    useEffect(() => {
        fetch("http://localhost:8080/drivers")
            .then(res => {
                if (!res.ok) throw new Error(`Server error: ${res.status}`)
                return res.json()
            })
            .then(data => setDrivers(Array.isArray(data) ? data : []))
            .catch((err) => {
                setError("Failed to load drivers: " + err.message)
                setDrivers([])
            })
    }, [])

    const handlerFilterChange = (event) => setFilterText(event.target.value)

    const filteredDrivers = drivers.filter(driver =>
        driver.username.toLowerCase().includes(filterText.toLowerCase()) ||
        driver.id.toString().includes(filterText) ||
        driver.status.toLowerCase().includes(filterText.toLowerCase())
    )

    const handleOpenPopup = (setter, resetFn) => {
        setError(null)
        resetFn()
        setter(true)
    }

    // Change Status
    const handleSubmitStatus = async () => {
        if (!statusForm.id || !statusForm.status) return setError("Please fill in all fields.")
        try {
            setIsLoading(true)
            setError(null)
            let response = await fetch(`http://localhost:8080/drivers/${statusForm.id}/status`, {
                method: "PATCH",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({status: statusForm.status}),
                credentials: "include"
            })
            if (response.status === 401) {
                await refreshAccessToken()
                const response = await fetch(`http://localhost:8080/drivers/${statusForm.id}/status`, {
                    method: "PATCH",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify({status: statusForm.status}),
                    credentials: "include"
                })
            }
            if (response.ok) {
                setOpenPopupChangeStatus(false)
                setDrivers(prev => prev.map(driver =>
                    driver.id.toString() === statusForm.id.toString()
                        ? {...driver, status: statusForm.status}
                        : driver
                ))
                setStatusForm({id: "", status: "driving"})
            }
        } catch {
            setError("Network error. Please try again.")
        } finally {
            setIsLoading(false)
        }
    }

    // Change Name
    const handleSubmitName = async () => {
        if (!nameForm.id || !nameForm.username) return setError("Please fill in all fields.")
        try {
            setIsLoading(true)
            setError(null)
            let response = await fetch(`http://localhost:8080/drivers/${nameForm.id}/username`, {
                method: "PATCH",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({username: nameForm.username}),
                credentials: "include"
            })
            if (response.status === 401) {
                await refreshAccessToken()
                response = await fetch(`http://localhost:8080/drivers/${nameForm.id}/username`, {
                    method: "PATCH",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify({username: nameForm.username}),
                    credentials: "include"
                })
            }
            if (response.ok) {
                setOpenPopupChangeName(false)
                setDrivers(prev => prev.map(driver =>
                    driver.id.toString() === nameForm.id.toString()
                        ? {...driver, username: nameForm.username}
                        : driver
                ))
                setNameForm({id: "", username: ""})
            }
        } catch {
            setError("Network error. Please try again.")
        } finally {
            setIsLoading(false)
        }
    }

    // Delete Driver
    const handleSubmitDelete = async () => {
        if (!deleteForm.id) return setError("Please enter a driver ID.")
        try {
            setIsLoading(true)
            setError(null)
            let response = await fetch(`http://localhost:8080/drivers/${deleteForm.id}`, {
                method: "DELETE",
                credentials: "include",
            })
            console.log("status:", response.status)
            console.log("ok:", response.ok)
            if (response.status === 401) {
                await refreshAccessToken()
                response = await fetch(`http://localhost:8080/drivers/${deleteForm.id}`, {
                    method: "DELETE",
                    credentials: "include",
                })
            }
            if (response.ok) {
                setOpenPopupDeleteDriver(false)
                setDrivers(prev => prev.filter(driver => driver.id.toString() !== deleteForm.id.toString()))
                setDeleteForm({id: ""})
            }
        } catch (err) {
            console.log("catch error:", err)
            setError("Network error. Please try again.")
        } finally {
            setIsLoading(false)
        }
    }

    return (
        <div className="drivers-table">
            <div className="drivers-input">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640" width="16" height="16" fill="#888">
                    <path
                        d="M544 513L397.2 364.2C417.2 336.3 429.1 302 429.1 265C429.1 171.9 354.4 96.1 262.6 96.1C170.7 96 96 171.8 96 264.9C96 358 170.7 433.8 262.5 433.8C302.3 433.8 338.8 419.6 367.5 395.9L513.5 544L544 513zM262.5 394.8C191.9 394.8 134.4 336.5 134.4 264.9C134.4 193.3 191.9 135 262.5 135C333.1 135 390.6 193.3 390.6 264.9C390.6 336.5 333.2 394.8 262.5 394.8z"/>
                </svg>
                <input value={filterText} onChange={handlerFilterChange} placeholder="Search"/>
                <div className="drivers-data-buttons">
                    <button onClick={() => handleOpenPopup(setOpenPopupChangeStatus, () => setStatusForm({
                        id: "",
                        status: "driving"
                    }))}>
                        Change Status
                    </button>
                    <button onClick={() => handleOpenPopup(setOpenPopupChangeName, () => setNameForm({
                        id: "",
                        username: ""
                    }))}>
                        Change Driver Name
                    </button>
                    <button onClick={() => handleOpenPopup(setOpenPopupDeleteDriver, () => setDeleteForm({id: ""}))}>
                        Delete Driver
                    </button>
                </div>
            </div>

            {openPopupChangeStatus &&
                <div className="driver-register">
                    <h2>Change Driver Status</h2>
                    <p>Driver ID</p>
                    <input
                        name="id"
                        value={statusForm.id}
                        onChange={e => setStatusForm({...statusForm, id: e.target.value})}
                        disabled={isLoading}
                        placeholder="Driver ID"
                    />
                    <p>Status</p>
                    <div className="status-options">
                        {["driving", "searching", "offline"].map(option => (
                            <button
                                key={option}
                                type="button"
                                className={`status-option ${statusForm.status === option ? "active" : ""}`}
                                onClick={() => setStatusForm({...statusForm, status: option})}
                                disabled={isLoading}
                            >
                                {option}
                            </button>
                        ))}
                    </div>
                    {error && <p className="register-error">{error}</p>}
                    <div className="drivers-register-buttons">
                        <button className="add-driver-cancel" onClick={() => setOpenPopupChangeStatus(false)}
                                disabled={isLoading}>Cancel
                        </button>
                        <button className="add-driver-save" onClick={handleSubmitStatus}
                                disabled={isLoading}>{isLoading ? "Saving..." : "Save"}</button>
                    </div>
                </div>
            }

            {openPopupChangeName &&
                <div className="driver-register">
                    <h2>Change Driver Name</h2>
                    <p>Driver ID</p>
                    <input
                        name="id"
                        value={nameForm.id}
                        onChange={e => setNameForm({...nameForm, id: e.target.value})}
                        disabled={isLoading}
                        placeholder="Driver ID"
                    />
                    <p>New Username</p>
                    <input
                        name="username"
                        value={nameForm.username}
                        onChange={e => setNameForm({...nameForm, username: e.target.value})}
                        disabled={isLoading}
                        placeholder="New Username"
                    />
                    {error && <p className="register-error">{error}</p>}
                    <div className="drivers-register-buttons">
                        <button className="add-driver-cancel" onClick={() => setOpenPopupChangeName(false)}
                                disabled={isLoading}>Cancel
                        </button>
                        <button className="add-driver-save" onClick={handleSubmitName}
                                disabled={isLoading}>{isLoading ? "Saving..." : "Save"}</button>
                    </div>
                </div>
            }

            {openPopupDeleteDriver &&
                <div className="driver-register">
                    <h2>Delete Driver</h2>
                    <h4 className="red-text">You also will delete<br/>
                        orders with this id</h4>
                    <p>Driver ID</p>
                    <input
                        name="id"
                        value={deleteForm.id}
                        onChange={e => setDeleteForm({id: e.target.value})}
                        disabled={isLoading}
                        placeholder="Driver ID"
                    />
                    {error && <p className="register-error">{error}</p>}
                    <div className="drivers-register-buttons">
                        <button className="add-driver-cancel" onClick={() => setOpenPopupDeleteDriver(false)}
                                disabled={isLoading}>Cancel
                        </button>
                        <button className="add-driver-save add-driver-delete" onClick={handleSubmitDelete}
                                disabled={isLoading}>{isLoading ? "Deleting..." : "Delete"}</button>
                    </div>
                </div>
            }

            <table>
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Username</th>
                    <th>Status</th>
                </tr>
                </thead>
                <tbody>
                {filteredDrivers.toReversed().map(driver => (
                    <tr key={driver.id}>
                        <td>{driver.id}</td>
                        <td>{driver.username}</td>
                        <td>{driver.status}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    )
}

export default AllDriversTable