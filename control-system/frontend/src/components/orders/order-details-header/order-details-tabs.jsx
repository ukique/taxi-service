import "./order-details-tabs.css"
import {useState} from "react";
import {useParams, useNavigate} from "react-router-dom";

function OrderDetailsTabs({ defaultTab }) {
    const [activeTab, setActiveTab] = useState(defaultTab)
    const navigate = useNavigate()
    const {id} = useParams();
    return (
        <div className="tabs">
            <button className={`tab ${activeTab === 'info' ? 'active' : ''}`}
                    onClick={() =>{ setActiveTab('info'); navigate(`/orders/page/${id}`)} }>
                Order info
            </button>
            <button className={`tab ${activeTab === 'location' ? 'active' : ''}`}
                    onClick={() =>{ setActiveTab('location'); navigate(`/orders/page/${id}/location`) }}>
                Location history
            </button>
        </div>
    )
}

export default OrderDetailsTabs