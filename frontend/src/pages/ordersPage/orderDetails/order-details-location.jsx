import "./order-details-location.css"
import "./order-details-info.css"
import OrderDetailsHeader from "../../../components/orders/order-details-header/order-details-header.jsx";
import { useParams } from "react-router-dom";
import Header from "../../../components/header/header.jsx";
import { useEffect, useState } from "react";

function OrderDetailsLocation() {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);
    const { id } = useParams();

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/location/${id}`);
                const json = await response.json();
                setData(json ?? []);
            } catch (err) {
                console.error("Failed to fetch location data:", err);
            } finally {
                setLoading(false);
            }
        };

        fetchData().catch(console.error);
    }, [id]);

    return (
        <>
            <Header />
            <OrderDetailsHeader defaultTab="location" />
            <div className="order-details-location">
                <div className="order-details-location-data">
                    <h1>Order Location History</h1>
                    <h2>Track and review the full<br />
                        route coordinates history</h2>
                </div>
                <div className="order-details-table">
                    <table>
                        <thead>
                        <tr>
                            <th>Driver ID</th>
                            <th>Order ID</th>
                            <th>Order Status</th>
                            <th>Lat</th>
                            <th>Lon</th>
                        </tr>
                        </thead>
                        <tbody>
                        {loading || data.length === 0 ? (
                            <tr>
                                <td colSpan={5} style={{ textAlign: "center", padding: "2rem", color: "#888" }}>
                                    {loading ? "Loading location history..." : "No location data found for this order"}
                                </td>
                            </tr>
                        ) : (
                            data.map((entry, index) => (
                                <tr key={index}>
                                    <td>{entry.driver_id}</td>
                                    <td>{entry.id}</td>
                                    <td>{entry.status}</td>
                                    <td>{entry.lat}</td>
                                    <td>{entry.lon}</td>
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

export default OrderDetailsLocation;