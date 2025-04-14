### GO + ECHO + GORM

CRUD

PORT = 1323

ENDPOINTS

- Products
  - [GET] /products
  - [GET] /products/{ID}
  - [POST] /products
  - [PUT] /products/{ID}
  - [DELETE] /products/{ID}

- Categories
    - [GET] /categories
    - [GET] /categories/{ID}
    - [POST] /categories
    - [PUT] /categories/{ID}
    - [DELETE] /categories/{ID}

- Carts
    - [GET] /carts
    - [GET] /carts/{ID}
    - [POST] /carts
    - [DELETE] /carts/{ID}
    - Items in cart
      - [POST] /carts/{cart_id}/{product_id}/{quantity}
      - [PUT] /carts/{cart_id}/{product_id}/{quantity}
    