package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yashbalyan08/system/routers"
)

func main() {
	//os.Setenv("MYSQL_PASSWORD", "mallusethono.8")
	fmt.Println("Backend")
	r := routers.Router()
	fmt.Println("Sever is starting...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening to server at 4000")
}
