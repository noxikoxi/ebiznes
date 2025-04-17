import {ReactNode} from "react";
import Navbar from "../components/Navbar.tsx";

type Props = {
    children : ReactNode
}
const Layout = ({children} : Props) => {
    return (
        <div id="main">
            <Navbar links={["/","/products", "/cart"]} labels={["main", "products", "cart"]}/>
            <div id="fill">
                {children}
            </div>
        </div>
    )
}

export default Layout;