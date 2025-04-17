export interface Product {
    id: number;
    name: string;
    price: number;
    description?: string;
}

export interface CartItem extends Product{
    quantity: number;
}

export interface RequestItem{
    product_id: number;
    quantity: number;
}

export interface RequestPayment{
    Items: RequestItem[];
    Total: number;
}
