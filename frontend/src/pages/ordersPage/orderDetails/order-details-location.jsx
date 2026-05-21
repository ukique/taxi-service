import "./order-details-location.css"
import "./order-details-info.css"
import OrderDetailsHeader from "../../../components/orders/order-details-header/order-details-header.jsx";
import {useParams} from "react-router-dom";
import Header from "../../../components/header/header.jsx";

function OrderDetailsLocation() {
    const {id} = useParams();
    return (
        <>
            <Header/>
            <OrderDetailsHeader defaultTab="location"/>
            <div className="order-details-location">
                <div className="order-details-location-data">
                    <h1>Order Location History</h1>
                    <h2>Track and review the full<br/>
                        route coordinates history </h2>
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
                        <tr>
                            <td>13</td>
                            <td>12</td>
                            <td>in_progress</td>
                            <td>-48.238709904837606</td>
                            <td>-48.238709904837606</td>
                        </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </>
    )
}

export default OrderDetailsLocation