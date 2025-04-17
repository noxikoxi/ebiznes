import {useEffect, useState} from "react";
import { useCartContext } from "../contexts/CartContext";
import {CartItem} from "../types/types.tsx";

const Cart = () => {
    const { cartItems , setCartItems} = useCartContext();
    const [totalPrice, setTotalPrice] = useState<number>(0);

    useEffect(() => {
        setTotalPrice(cartItems.reduce(
            (sum, item) => sum + item.price * item.quantity,
            0
        ))
    }, [cartItems]);

    const addQuantity = (item: CartItem) => {
        const updatedItems = cartItems.map(cartItem => {
            if (cartItem.id === item.id) {
                return { ...cartItem, quantity: cartItem.quantity + 1 };
            }
            return cartItem;
        });
        setCartItems(updatedItems);
    }

    const lowerQuantity = (item: CartItem) => {
        const updatedItems = cartItems.map(cartItem => {
            if (cartItem.id === item.id) {
                return { ...cartItem, quantity: cartItem.quantity - 1 };
            }
            return cartItem;
        }).filter((item) => item.quantity !== 0);
        setCartItems(updatedItems);
    }

    return (
        <div>
            <h2>Your Cart</h2>
            {cartItems.length === 0 ? (
                <p>Your cart is empty</p>
            ) : (
                <div>
                    <ul>
                        {cartItems.map((item) => (
                            <li key={item.id}>
                                <div style={{display: "flex", gap: 10, placeContent: "space-between"}}>
                                    <strong>{item.name}</strong>
                                    <div style={{display: "flex", gap: 2}}>
                                        <button className="add cartButton" onClick={() => addQuantity(item)}>+</button>
                                        <button className="danger cartButton" onClick={() => lowerQuantity(item)}>-</button>
                                    </div>
                                </div>
                                <span>
                                {item.quantity} ×{" "}
                                {item.price.toFixed(2)} zł ={" "}
                                {(item.price * item.quantity).toFixed(2)} zł
                                    </span>
                            </li>
                        ))}
                    </ul>

                    <h3>Total: {totalPrice.toFixed(2)} zł</h3>
                </div>
            )}
        </div>
    );
};

export default Cart;
