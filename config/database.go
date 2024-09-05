package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

//const connectionString string = "root:mallusethno.8@/transport"

var db *sql.DB

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
		Passwd: "mallusethno.8",
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
