package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var MYSQL_PASSWORD string = os.Getenv("MYSQL")

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// pass := os.Getenv("MYSQL_PASSWORD")
	// fmt.Println(pass)
	cfg := mysql.Config{
		User:   "root",
		Passwd: MYSQL_PASSWORD,
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "transport",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	checkNilErr(err)
	fmt.Println("The db: ", db)
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("The server is connected")
}

func GetDB() *sql.DB {
	return db
}
