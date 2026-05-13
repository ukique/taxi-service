import "./allOrdersTable.css"
import {useCallback, useState} from "react";
import {useSubscription} from "../../hooks/useSubscription";

function AllOrdersTable() {
    const [orders, setOrders] = useState([]);
    const [filterText, setFilterText] = useState("");

    const handleMessage = useCallback((data) => {
        if (data.type === "orders") {
            setOrders(data.data ?? []);
        }
    }, []);

    useSubscription({
        subscribeMsg: {type: "subscribe_orders", page: 1},
        onMessage: handleMessage,
    });

    const handlerFilterChange = (event) => setFilterText(event.target.value);

    const filteredOrders = orders.filter(order =>
        String(order.id).includes(filterText) ||
        (order.status ?? "").toLowerCase().includes(filterText.toLowerCase())
    );

    return (
        <>
            <div className="drivers-input">
                <div className="orders-table">
                    <input value={filterText} onChange={handlerFilterChange} placeholder="Search"/>
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
                        {filteredOrders.length === 0 ? (
                            <tr>
                                <td colSpan={4} style={{textAlign: "center", padding: "2rem", color: "#888"}}>
                                    {orders.length === 0 ? "Loading orders..." : "No orders match your search"}
                                </td>
                            </tr>
                        ) : (
                            filteredOrders.toReversed().map(order => (
                                <tr key={order.id}>
                                    <td><a className="order-status-id" href={`/orders/page/${order.id}`}>{order.id}</a>
                                    </td>
                                    <td>{order.driver_id}</td>
                                    <td>{order.status}</td>
                                    <td>{order.created_at}</td>
                                </tr>
                            ))
                        )}
                        </tbody>
                    </table>
                </div>
            </div>
        </>
    );
}

export default AllOrdersTable;