package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	database "github.com/yashbalyan08/system/config"
)

func checkNilErr(err error) {
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		log.Fatal(err)
	}
}

var db *sql.DB = database.GetDB()

func checkLogin(params map[string]string) bool {
	query := "select password from users where id = ?"
	result, err := db.Prepare(query)
	checkNilErr(err)

	defer result.Close()

	var pass string
	err = result.QueryRow(params["id"]).Scan(&pass)

	checkNilErr(err)

	return pass == params["pass"]
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)

	if checkLogin(params) {
		json.NewEncoder(w).Encode("Logged in successful")
		return
	} else {
		json.NewEncoder(w).Encode("Not Logged in")
	}
}
