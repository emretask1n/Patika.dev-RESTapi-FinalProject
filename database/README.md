# Database

![Database Schema](DbSchema.png)


The database schema I use for the project is here. In order for you to try the API, You can execute the **DbSchema.sql** file in mySQL.

To connect to the database you need to configure the code in **database.go**;

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
