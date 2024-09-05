package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yashbalyan08/system/routers"
)

func main() {
	fmt.Println("Backend")
	r := routers.Router()
	fmt.Println("Sever is starting...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening to server at 4000")
}
