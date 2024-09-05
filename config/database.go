package db

import (
	"context"
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

func dbConnect(ctx context.Context) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "mallusethno.8",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
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

func init() {
	ctx := context.TODO()
	dbConnect(ctx)
}

func GetDB() *sql.DB {
	defer db.Close()
	return db
}
