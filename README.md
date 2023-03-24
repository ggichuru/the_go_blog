# the_go_blog

Learning GO by building a CRUD restful API with gin. gorm and postgresql

## we are also going to learn how to

+ build models in go
+ connect to the running SQL server
+ run db migrations with GORM

---

> **What is GORM ?**
>
> + [GORM](https://gorm.io/) is an object-relational mapper (ORM) library for Golang, buit on top of the golang [database/sql](https://pkg.go.dev/database/sql) package.
> + It only works with relational Databases (PostrgreSQL, SQLite, MySQL)
> + we will use this to create db models
> + Provides a way to manage conn b2n Go structs and corresponding SQL reps in the DB
> + GORM model is a struct with basic golang types, pointers and custom types implementing *Scanner* and *Valuer*.
> + Use `PascalCase` syntax for model and column name and use `snake_case` for columns.
> + Stuct tags are optional
> + Gorm allows you to use the native DB types in the `gorm:""` tag annotation.
>
> ---
>
> **What is Gin Gonic ?**
>
> + [Gin Gonic](https://gin-gonic.com/docs/) is a high-perfomant web framework written in go.
> + Same class with the likes of express
