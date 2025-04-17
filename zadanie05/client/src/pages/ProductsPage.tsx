import {useEffect} from "react";
import Products from "../components/Products.tsx";

const ProductsPage = () => {
    useEffect(() => {
        document.title = 'Products'
    }, [])

    return (
        <div className="f1 centered">
            <Products/>
        </div>
    )
}

export default ProductsPage;