import "./resizable.css"
import pinIcon from "../../assets/maps-and-flags.svg"

function Resizable(){
    return(
        <div className="resizable-card">
            <div className="resizable-input-wrapper">
                <div className="resizable-pin">
                    <img src={pinIcon} alt="pin" width={16} height={16} />
                </div>
                <input className="resizable-input" placeholder="Add a pick-up location" />
            </div>
            <div className="resizable-input-wrapper">
                <div className="resizable-pin">
                    <img src={pinIcon} alt="pin" width={16} height={16} />
                </div>
                <input className="resizable-input" placeholder="Add your destination" />
            </div>
            <button className="resizable-order">
                Order
            </button>
        </div>
    )
}
export default Resizable