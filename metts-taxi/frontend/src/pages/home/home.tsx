import "./home.css"
import HomeMenu from "../../components/home-menu/home.tsx";
import Resizable from "../../components/home-resizable/resizable.tsx";


function Home() {
    return (
        <>
        <HomeMenu/>
        <div className="home-page">
        <p className="home-headline">Where do you want to go?</p>
        <Resizable/>
        </div>
        </>
    )
}

export default Home