import "./allOrdersTable.css"
import { useEffect, useState } from "react";
import { useWS } from '../../utils/useWS';

function AllOrdersTable() {
    const [orders, setOrders] = useState([]);
    const [filterText, setFilterText] = useState("");
    const { sendMessage, onMessage, isConnected } = useWS();

    useEffect(() => {
        if (isConnected) {
            sendMessage({ type: 'orders', pageID: 1 });
        }
    }, [isConnected, sendMessage]);

// 2. add onMessage
    useEffect(() => {
        onMessage((data) => {
            if (data.type === 'orders') {
                setOrders(data.data);
            }
        });
    }, [onMessage]);

    const handlerFilterChange = (event) => setFilterText(event.target.value);
    const filteredOrders = orders.filter(order =>
        String(order.id).includes(filterText) ||
        (order.status ?? "").toLowerCase().includes(filterText.toLowerCase())
    );

    return (
        <>
            <div className="drivers-input">
                <div className="drivers-table">
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
    );
}

export default AllOrdersTable;