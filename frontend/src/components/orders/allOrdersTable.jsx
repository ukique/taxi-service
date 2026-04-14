import "./allOrdersTable.css"
import {useEffect, useState} from "react";

function AllOrdersTable() {
    const [error, setError] = useState(null)

    //Order Table
    const [orders, setOrders] = useState([])
    useEffect(() => {
        fetch("http://localhost:8080/orders")
            .then(res => {
                if (!res.ok) throw new Error(`Server error: ${res.status}`)
                return res.json()
            })
            .then(data => setOrders(Array.isArray(data) ? data : []))
            .catch((err) => {
                setError("Failed to load drivers: " + err.message)
                setOrders([])
            })
    }, [])


    //filter
    const [filterText, setFilterText] = useState("")
    const handlerFilterChange = (event) => setFilterText(event.target.value)
    const filteredOrders = orders.filter(order =>
        String(order.id).includes(filterText) ||
        (order.status ?? "").toLowerCase().includes(filterText.toLowerCase())
    )
    return (
        <>
            <div className="drivers-input">
                <div className="drivers-table">
                    <input value={filterText} onChange={handlerFilterChange} placeholder="Search"/>
                    {error}
                    <table>
                        <thead>
                        <tr>
                            <th>ID</th>
                            <th>Driver ID</th>
                            <th>Status</th>
                            <th>Created_At</th>
                        </tr>
                        </thead>
                        <tbody>
                        {filteredOrders.toReversed().map(order => (
                            <tr key={order.id}>
                                <td><a className="order-status-id" href={`/orders/page/${order.id}`}>{order.id}</a></td>
                                <td>{order.driver_id}</td>
                                <td>{order.status}</td>
                                <td>{order.created_at}</td>
                            </tr>
                        ))}
                        </tbody>
                    </table>
                </div>
            </div>
        </>
    )
}

export default AllOrdersTable