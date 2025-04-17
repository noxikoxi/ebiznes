import {useEffect} from "react";

const MainPage = () => {
    useEffect(() => {
        document.title = 'Main Page'
    }, [])

    return (
        <div className="centered col">
            <h2>REACT + GO</h2>
            <h2>Products, Carts, Payments</h2>
            <p>See all Products</p>
            <p>Add Products to cart</p>
            <p>Pay</p>
        </div>
    )
}

export default MainPage;