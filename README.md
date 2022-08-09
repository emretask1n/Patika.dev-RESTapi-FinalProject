# Patika.dev & Property Finder Final Project

## REST API service in Golang

A REST API for shopping cart system that allows users to  List Products, Add Item To Cart, Show Cart Delete Cart items, and Complete Orders.

### Frameworks Used For the Project

- GORM
- Fiber

### Relational Database

- MySQL

### Go Requirements
execute followings
```
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/fiber/v2/middleware/cors
go get gorm.io/gorm
go get gorm.io/driver/mysql
```


### Endpoints

| Method |               URL               |                                              Function |
|--------|:-------------------------------:|------------------------------------------------------:|
| GET    |            /products            |                                         Show Products |
| GET    |         /cart/:user_id          |                                     Show cart of user |
| GET    |          /givenAmount           |                                     Show Given Amount |
| DELETE |  /delete/:user_id/:product_id   |     Delete item from cart with user id and product id |
| POST   | /:product_id/:user_id/:quantity | Add item to cart with product id user id and quantity |
| POST   |         /order/:user_id         |                           Complete order with user id |
| POST   |         /amount/:amount         |                              Set a new "Given Amount" |



### Database

The database schema I use for the project is in the database file. In order for you to try the API, You can execute the **DbSchema.sql** file in mySQL.

To connect to the database you need to configure the code in **database/database.go**;

```
dsn := "root:1234@tcp(127.0.0.1:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"
```

A DSN in its fullest form:

```
username:password@protocol(address)/dbname?param=value
```

If you have problems with data insertion, you can use the codes I added comments in the  file.

```
SET SQL_SAFE_UPDATES = 0;

SET FOREIGN_KEY_CHECKS = 0;
```

To enable it again, just type 1 instead of 0 and run it.
