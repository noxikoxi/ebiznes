#### Endpointy i struktura przyjmowanych json

##### Products

- [GET] products
- [GET] products/{id}
- [POST] products
``` 
{
    "name": "Telewizor",
    "price": 5000.0,
    "category_id": 1
}
```

- [PUT] products/{id}
```  
{
    "name": "telwizor",
    "price": 9999.0,
    "category_id": 1
}
```
- [DELETE] products/{id}

##### Categories

- [GET] categories
- [GET] categories/{id}
- [POST] categories
``` 
{
    "name": "Żywność"
}
```

- [PUT] categories/{id}
 ```
{
    "name": "niekoniecznie żywność"
}
```
- [DELETE] categories/{id}

##### Carts

- [GET] carts
- [GET] carts/{id}
- [POST] carts
```
{
    "products": [
        {
            "product_id": 1,
            "quantity": 10
        },
        {
            "product_id": 5,
            "quantity": 20
        },
        {
            "product_id": 6,
            "quantity": 2
        }
    ]

}
```
- [PUT] carts/{id}
```
{
    "product_id": 2,
    "quantity": 99
}
```
- [DELETE] carts/{id}





