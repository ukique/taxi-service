import "./resizable.css"
import pinIcon from "../../assets/maps-and-flags.svg"
import { useState } from "react";

interface LatLng {
    lat: number;
    lng: number;
}

function randomLatLng(): LatLng {
    const lat = Math.random() * 180 - 90;
    const lng = Math.random() * 360 - 180;
    return { lat, lng };
}

function Resizable() {
    const [pickup, setPickup] = useState("");
    const [destination, setDestination] = useState("");

    const handlePinClick = (target: "pickup" | "destination") => {
        const { lat, lng } = randomLatLng();
        const coords = `${lat.toFixed(6)}, ${lng.toFixed(6)}`;

        if (target === "pickup") {
            setPickup(coords);
        } else {
            setDestination(coords);
        }
    };

    return (
        <div className="resizable-card">
            <div className="resizable-input-wrapper">
                <div
                    className="resizable-pin"
                    onClick={() => handlePinClick("pickup")}
                    style={{ cursor: "pointer" }}
                >
                    <img src={pinIcon} alt="pin" width={16} height={16} />
                </div>
                <input
                    className="resizable-input"
                    placeholder="Add a pick-up location"
                    value={pickup}
                    onChange={(e) => setPickup(e.target.value)}
                />
            </div>
            <div className="resizable-input-wrapper">
                <div
                    className="resizable-pin"
                    onClick={() => handlePinClick("destination")}
                    style={{ cursor: "pointer" }}
                >
                    <img src={pinIcon} alt="pin" width={16} height={16} />
                </div>
                <input
                    className="resizable-input"
                    placeholder="Add your destination"
                    value={destination}
                    onChange={(e) => setDestination(e.target.value)}
                />
            </div>
            <button className="resizable-order">
                Order
            </button>
        </div>
    )
}

export default Resizable