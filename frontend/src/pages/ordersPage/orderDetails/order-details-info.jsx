import OrderDetailsHeader from "../../../components/orders/order-details-header/order-details-header.jsx";
import "./order-details-info.css"
import Header from "../../../components/header/header.jsx";

function OrderDetailsInfo() {

    return (
        <>
            <Header/>
            <OrderDetailsHeader defaultTab="info"/>
            <div className="order-details-info">
                <div className="current-location">
                    <h1 id="current-location-info">Current Location</h1>
                    <div className="current-location-data">
                        <h1>Lat: </h1>
                        <h2>-48.238709904837606</h2>
                    </div>
                    <div className="current-location-data">
                    <h1>Lon: </h1>
                    <h2>-48.238709904837606</h2>
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
                            <tr>
                                <td>1</td>
                                <td>1</td>
                                <td>in_progress</td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </>
    )
}

export default OrderDetailsInfo