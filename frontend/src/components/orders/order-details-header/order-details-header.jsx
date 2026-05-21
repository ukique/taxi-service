import Header from "../../header/header.jsx";
import "./order-details-header.css"
import {Link, useParams} from "react-router-dom";
import backArrow from '../../../assets/backArrow.png';
import OrderDetailsTabs from "./order-details-tabs.jsx";

function OrderDetailsHeader({ defaultTab }) {
    const {id} = useParams();
    return (
        <>
            <div className="orders-details-header">
                <div className="back-to-orders">
                    <Link to="/orders">
                        <img src={backArrow} alt="back"/>
                        <h2>Back to orders</h2>
                    </Link>
                </div>
                <div className="orders-details-header-data">
                    <h1>Order Details</h1>
                    <h2>#{id}</h2>
                </div>
                <OrderDetailsTabs defaultTab={defaultTab} />
            </div>
         </>
    )
}

export default OrderDetailsHeader