import OrderDetailsHeader from "../../../components/orders/order-details-header/order-details-header.jsx";
import "./order-details-info.css"
import Header from "../../../components/header/header.jsx";
import {useCallback, useState} from "react";
import {useSubscription} from "../../../hooks/useSubscription";
import {useParams} from "react-router-dom";

function OrderDetailsInfo() {
    const [coordinates, setCoordinates] = useState(null);
    const {id} = useParams();

    const handleMessage = useCallback((data) => {
        if (data.type === "coordinates") {
            setCoordinates(data.data);
        }
    }, []);

    useSubscription({
        subscribeMsg: {type: "subscribe_orderDetails", page: Number(id)},
        onMessage: handleMessage,
    });

    return (
        <>
            <Header/>
            <OrderDetailsHeader defaultTab="info"/>
            <div className="order-details-info">
                <div className="current-location">
                    <h1 id="current-location-info">Current Location</h1>
                    <div className="current-location-data">
                        <h1>Lat: </h1>
                        <h2>{coordinates ? coordinates.lat : "—"}</h2>
                    </div>
                    <div className="current-location-data">
                        <h1>Lon: </h1>
                        <h2>{coordinates ? coordinates.lon : "—"}</h2>
                    </div>
                </div>
                <div className="order-details-data">
                    <h1 id="order-details-h1">Order Details</h1>
                    <div className="order-details-table">
                        <table>
                            <thead>
                            <tr>
                                <th>Driver ID</th>
                                <th>Order ID</th>
                                <th>Order Status</th>
                            </tr>
                            </thead>
                            <tbody>
                            {coordinates ? (
                                <tr>
                                    <td>{coordinates.driver_id}</td>
                                    <td>{coordinates.id}</td>
                                    <td>{coordinates.status || "—"}</td>
                                </tr>
                            ) : (
                                <tr>
                                    <td colSpan={3} style={{textAlign: "center", padding: "2rem", color: "#888"}}>
                                        Waiting for order data...
                                    </td>
                                </tr>
                            )}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </>
    )
}

export default OrderDetailsInfo