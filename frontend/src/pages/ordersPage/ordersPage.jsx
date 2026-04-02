import "./ordersPage.css"
import Header from "../../components/header/header.jsx";
import CreateOrderButton from "../../components/orders/createOrder.jsx";
import AllOrdersTable from "../../components/orders/allOrdersTable.jsx";

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
                <AllOrdersTable/>
            </div>
        </>
    )
}

export default Orders