import "./ordersPage.css"
import Header from "../../components/header/header.jsx";
import CreateOrderButton from "../../components/orders/createOrder.jsx";

function Orders() {
    return (
        <>
            <div className="orders">
                <Header/>
                <div className="orders-main">
                    <h1>Orders</h1>
                    <CreateOrderButton />
                </div>
                <h4>Manage and track all active rides</h4>
            </div>
        </>
    )
}

export default Orders