package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	// data source name
	dsn := "root:12345678@tcp(127.0.0.1)/go_test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
