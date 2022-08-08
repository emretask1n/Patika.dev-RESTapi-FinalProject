# Patika.dev & Property Finder Final Project

## REST API service in Golang

### Frameworks Used For the Project

- GORM
- Fiber

### Relational Database

- MySQL


### Endpoints

| Method |                     URL                      |              Function |
|--------|:--------------------------------------------:|----------------------:|
| GET    |                  /products                   |         Show Products |
| GET    |                /cart/:user_id                |             Show cart |
| GET    |                 /givenAmount                 |     Show Given Amount |
| DELETE |         /delete/:user_id/:product_id         | Delete item from cart |
| POST   |       /:product_id/:user_id/:quantity        |      Add item to cart |
| POST   |                  /order/:id                  |         CompleteOrder |
| POST   |           /amount/:amount                    |  Set new given amount |



### Database

The database schema I use for the project is in the database file. In order for you to try the API, the sql codes that insert data into the Product and User tables are also in the same file path.

To connect to the database you need to configure the code;

```go
dsn := "root:1234@tcp(127.0.0.1:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"
```

A DSN in its fullest form:

```go
username:password@protocol(address)/dbname?param=value
```

If you have problems with data insertion, you can use the codes I added comments in the **DataInsertion.sql** file.

```sql
SET SQL_SAFE_UPDATES = 0;

SET FOREIGN_KEY_CHECKS = 0;
```

To enable it again, just type 1 instead of 0 and run it.
