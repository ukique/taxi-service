import Header from "../../header/header.jsx";
import "./orderDetails.css"
import {Link, useParams} from "react-router-dom";
import backArrow from '../../../assets/backArrow.png';
import OrderDetailsData from "./orderDetailsData.jsx";

function OrderDetails() {
    const {id} = useParams();
    return (
        <>
            <Header/>
            <div className="orders-details">
                <div className="back-to-orders">
                    <Link to="/orders">
                        <img src={backArrow} alt="back"/>
                        <h2>Back to orders</h2>
                    </Link>
                </div>
                <div className="orders-details-info">
                    <h1>Order Details</h1>
                    <h2>#{id}</h2>
                </div>
            <OrderDetailsData/>
            </div>
        </>
    )
}

export default OrderDetails