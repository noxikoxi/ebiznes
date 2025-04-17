import {useEffect} from "react";
import Cart from "../components/Cart.tsx";
import Payment from "../components/Payment.tsx";

const CartPage = () => {
    useEffect(() => {
        document.title = 'Cart'
    }, [])
    return (
        <div>
            <Cart/>
            <Payment/>
        </div>
    )
}

export default CartPage;