import "./drivers-history.css"
import {Link, useParams} from "react-router-dom";
import Header from "../../../components/header/header.jsx";
import {useEffect, useState} from "react";
import LowerHeader from "../../../components/header/lowerHeader.jsx";
import {refreshAccessToken} from "../../../api/authApi.js";

function DriversHistory() {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);
    const {id, pageID} = useParams();
    const page = Number(pageID);

    useEffect(() => {
        const fetchData = async () => {
            try {
                let response = await fetch(`http://localhost:8080/drivers/${id}/page/${pageID}`, {
                    credentials: "include",
                });
                if (response.status === 401) {
                    await refreshAccessToken();
                    response = await fetch(`http://localhost:8080/drivers/${id}/page/${pageID}`, {
                        credentials: "include",
                    });
                }
                const json = await response.json();
                setData(json ?? []);
            } catch (err) {
                console.error("Failed to fetch drivers history:", err);
            } finally {
                setLoading(false);
            }
        };

        fetchData().catch(console.error);
    }, [id, pageID]);

    return (
        <>
            <Header/>
            <div className="drivers-history-info">
            <h1>Driver #{id} History</h1>
            </div>
            <div className="drivers-history">
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
                    {loading || data.length === 0 ? (
                        <tr>
                            <td colSpan={5} style={{ textAlign: "center", padding: "2rem", color: "#888" }}>
                                {loading ? "Loading driver history..." : "No driver data found!"}
                            </td>
                        </tr>
                    ) : (
                        data.map((entry, index) => (
                            <tr key={index}>
                                <td>{entry.DriverID}</td>
                                <td>{entry.id}</td>
                                <td>{entry.status || "—"}</td>
                                <td>{entry.lat}</td>
                                <td>{entry.lon}</td>
                            </tr>
                        ))
                    )}
                    </tbody>
                </table>
            </div>
                <div className="drivers-pagination">
                    {page > 1 && (
                        <h3>
                            <Link to={`/drivers/${id}/page/${page - 1}`}>
                                {page - 1}
                            </Link>
                        </h3>
                    )}

                    <h2>
                        <Link to={`/drivers/${id}/page/${page}`}>
                            {page}
                        </Link>
                    </h2>

                    <h3>
                        <Link to={`/drivers/${id}/page/${page + 1}`}>
                            {page + 1}
                        </Link>
                    </h3>
                </div>
            <LowerHeader/>

            </div>
        </>
    )
}

export default DriversHistory