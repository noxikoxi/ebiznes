import { useState } from "react";
import { useCartContext } from "../contexts/CartContext";
import axios from "axios";
import {RequestPayment} from "../types/types.tsx";

const Payment = () => {
    const { cartItems, setCartItems } = useCartContext();
    const [isConfirming, setIsConfirming] = useState(false);

    const sendToServer = async () => {
        try {
            const request : RequestPayment = {
              Items: cartItems.map((item) => ({
                  product_id: item.id,
                  quantity: item.quantity,
              })),
                Total: cartItems.reduce((total, item) => total + item.price * item.quantity, 0),
            };


            await axios.post("http://localhost:1323/payment", request);
            setCartItems([]);
        } catch (error) {
            console.error("Error sending payment:", error);
        }
    };

    const handlePaymentClick = () => {
        if (cartItems.length !== 0) {
            setIsConfirming(true);
        }
    };

    const handleConfirmPayment = () => {
        sendToServer();
        setIsConfirming(false);
    };

    const handleCancelPayment = () => {
        setIsConfirming(false);
    };

    return (
        <div>
            <button onClick={handlePaymentClick}>Payment</button>

            {isConfirming && (
                <div className="modalBG">
                    <div className="modal">
                        <h3>Confirm Payment</h3>
                        <p>Are you sure you want to send your cart for payment?</p>
                        <div style={{display: "flex", gap: 10}}>
                            <button className="add" onClick={handleConfirmPayment}>Confirm</button>
                            <button className="danger" onClick={handleCancelPayment}>Cancel</button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Payment;
