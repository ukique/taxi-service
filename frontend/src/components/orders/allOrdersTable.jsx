import "./allOrdersTable.css"
import {useCallback, useEffect, useMemo, useRef, useState} from "react";
import {useSubscription} from "../../hooks/useSubscription";
import {Link, useParams} from "react-router-dom";
import LowerHeader from "../header/lowerHeader.jsx";
import {refreshAccessToken} from "../../api/authApi.js";

function AllOrdersTable() {
    const [orders, setOrders] = useState([]);
    const [filterText, setFilterText] = useState("");

    const {id} = useParams();
    const page = Number(id);

    const pageRef = useRef(page);
    useEffect(() => {
        pageRef.current = page;
    }, [page]);

    useEffect(() => {
        setOrders([]);
    }, [page]);

    const handleMessage = useCallback((data) => {
        if (data.type === "orders" && pageRef.current === 1) {
            setOrders(data.data ?? []);
        }
    }, []);

    const subscribeMsg = useMemo(() =>
            page === 1 ? {type: "subscribe_orders", page: 1} : null
        , [page]);

    useSubscription({
        subscribeMsg,
        onMessage: handleMessage,
    });

    useEffect(() => {
        if (page === 1) return;

        const fetchData = async () => {
            try {
                let response = await fetch(`http://localhost:8080/orders/${page}`, {
                    credentials: "include",
                });
                if (response.status === 401) {
                    await refreshAccessToken();
                    response = await fetch(`http://localhost:8080/orders/${page}`, {
                        credentials: "include",
                    });
                }
                const json = await response.json();
                setOrders(Array.isArray(json) ? json : []);  // защита если сервер вернул не массив
            } catch (err) {
                console.error("Failed to fetch orders:", err);
            }
        };

        fetchData().catch(console.error);
    }, [page]);

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
                                    <td>
                                        <Link className="link" to={`/orders/page/${order.id}`}>
                                            {order.id}
                                        </Link>
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
            <div className="orders-pagination">
                {page > 1 && (
                    <h3>
                        <Link to={`/orders/${page - 1}`}>
                            {page - 1}
                        </Link>
                    </h3>
                )}

                <h2>
                    <Link to={`/orders/${page}`}>
                        {page}
                    </Link>
                </h2>

                <h3>
                    <Link to={`/orders/${page + 1}`}>
                        {page + 1}
                    </Link>
                </h3>
                <LowerHeader/>
            </div>
        </>
    );
}

export default AllOrdersTable;