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
                            <th>User ID</th>
                            <th>Driver ID</th>
                            <th>PickUp_Lat</th>
                            <th>PickUp_Lon</th>
                            <th>DropOut_Lat</th>
                            <th>DropOut_Lon</th>
                            <th>Status</th>
                            <th>Created_At</th>
                            <th>Updated_At</th>
                        </tr>
                        </thead>
                        <tbody>
                        {filteredOrders.toReversed().map(order => (
                            <tr key={order.id}>
                                <td><a className="order-status-id" href={`/orders/page/${order.id}`}>{order.id}</a></td>
                                <td>{order.user_id}</td>
                                <td>{order.driver_id}</td>
                                <td>{order.pick_up_lat}</td>
                                <td>{order.pick_up_lon}</td>
                                <td>{order.drop_out_lat}</td>
                                <td>{order.drop_out_lon}</td>
                                <td>{order.status}</td>
                                <td>{order.created_at}</td>
                                <td>{order.updated_at}</td>
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