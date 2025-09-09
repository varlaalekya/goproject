package main

import (
	"apigolang/api"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_"github.com/go-sql-driver/mysql"
)

func main() {
	dns := "root:root@tcp(127.0.0.1:3306)/myfile?parseTime=true"
	
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Error opening database: ", err)

	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established successfully!")

	api.RegisterRoutes(db)

	log.Println("Server started on :3030")
	log.Fatal(http.ListenAndServe(":3030",nil))
}