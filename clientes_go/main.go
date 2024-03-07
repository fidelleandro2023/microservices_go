package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	//_ "github.com/microsoft/go-mssqldb"
)

// @Golang API REST
var server = `localhost\SQLEXPRESS`
var port = 1433
var user = "fidel"
var password = "1234"
var database = "VENTAS_PASAJES"
var db *sql.DB

const PORT string = ":8081"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "connected!!!")

	})
	log.Fatal(http.ListenAndServe(":8081", nil))
	//SelectVersion()
}
