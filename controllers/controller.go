package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	database "github.com/yashbalyan08/system/config"
)

var db *sql.DB = database.GetDB()

func RetrieveAvailableDriver(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RetrieveAvailableDriver")
}

func MakeDriverAvailable() {
	fmt.Println("MakeDriverAvailable")
}

func MakeDriverUnavailable() {
	fmt.Println("MakeDriverUnavailable")
}
