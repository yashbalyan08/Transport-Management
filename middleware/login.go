package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	database "github.com/yashbalyan08/system/config"
	user "github.com/yashbalyan08/system/models"
)

func checkNilErr(err error) {
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		log.Fatal(err)
	}
}

var db *sql.DB = database.GetDB()

func checkLogin(user user.User) bool {

	query := "select password from users where id = ?"
	result, err := db.Prepare(query)
	checkNilErr(err)

	// defer result.Close()
	var pass string
	err = result.QueryRow(user.Id).Scan(&pass)

	checkNilErr(err)

	return pass == user.Password
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user user.User

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Printf("Check params: %v\n", user)

	if checkLogin(user) {
		json.NewEncoder(w).Encode("Logged in successful")
		return
	} else {
		json.NewEncoder(w).Encode("Not Logged in")
	}
	db.Close()
}
