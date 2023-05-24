# Ecommerce with Go and mongo

This repository implements a Go API built with gin that imitates how a ecommerce would work. It allows you to signup and login as an user with jwt authentication, add documents to a mongo collection, both for users and for cart items. It allows you to generate product orders and to calculate the cart total with aggregation pipelines.

## Implements
Go
gin
mongo 

## Installation and configuration instructions

1. Clone the github repository 
2. Start the mongo container `docker-compose up -d`
3. Run the gin API `go run main.go`
4. Access to the path operations: `localhost:8000/{path}`

## Path Operations

This project has several path operations for both user and cart handler functions

- /signup: Allows you to receive a JSON object and insert the new user into a mongo collection.
```
{
  "first_name": "user name",
  "last_name": "user last name",
  "email": "useremail@gmail.com",
  "password": "password123",
  "phone": "5511856421"
}
```

- /login: Validates user credentials and generates a token for jwt authentication
```
{
  "email": "useremail@gmail.com",
  "password": "password123"
}
```

- /admin/addproduct: Add new objects into the products mongo collection
```
{
    "product_name": "Glass",
    "price": 50,
    "rating": 3,
    "image": "rglass.jpg"
}
```

- /users/productview: Returns a slice with all products stored in the mongo collection

- /users/search?name={name}: Performs a regular expression (regex) pattern matching, it receives a "name" query parameter

For the following path operations is required to provide the login token as a header

- /addtocart?id={productid}&userID={userID}: Add ProductUser into User model as a slice and update user collection

- /removeitem?id={productid}&userID={userID}: Remove ProductUser from User model and update user collection

- /listcart?id={userid}: List all products stored in user cart

- /cartcheckout?id={userid}: Places an order slice into the user model

