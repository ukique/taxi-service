import "./home.css"
import menuIcon from "../../assets/menu.svg"

function HomeMenu(){
    const userName = localStorage.getItem("username");
    return (
        <div className="home-menu">
            <span className="home-menu-username">{userName}</span>
            <button className="home-menu-button">
                <img alt="home-menu" src={menuIcon}/>
            </button>
        </div>
    )
}
export default HomeMenu