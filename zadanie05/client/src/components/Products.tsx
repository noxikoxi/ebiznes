import {useEffect, useState} from "react";
import {useCartContext} from "../contexts/CartContext.tsx";
import {Product} from "../types/types.tsx";
import axios from "axios";

const Products = () => {
    const { cartItems, setCartItems } = useCartContext();
    const [products, setProducts] = useState<Product[]>([]);

    useEffect(() => {
        console.log(cartItems);

        axios.get<Product[]>("http://localhost:1323/products")
            .then(res => {
                setProducts(res.data);
            })
            .catch(err => console.error(err));
    }, []);

    const addToCart = (product: Product) => {
        const existingItemIndex = cartItems.findIndex(
            (item) => item.id === product.id
        );
        if (existingItemIndex !== -1) {
            const updatedItems = [...cartItems];
            updatedItems[existingItemIndex].quantity += 1;
            setCartItems(updatedItems);
        } else {
            setCartItems([...cartItems, { ...product, quantity: 1 }]);
        }
    }

    return (
        <>
            <h2>All Products</h2>
            <div className="productsWrapper">
                {products.map((product) => (
                    <div key={product.id} className="productContainer">
                        <div className="productWrapper">
                            <div className="product">
                                <span>{product.name}</span>
                                <span>{product.description}</span>
                                <span>{product.price} zł</span>
                            </div>
                            <button onClick={() => addToCart(product)}>Add to cart</button>
                        </div>
                    </div>
                ))}
            </div>
        </>
    );
};


export default Products;