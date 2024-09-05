package routers

import (
	"github.com/gorilla/mux"
	middleware "github.com/yashbalyan08/system/middleware"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", middleware.CheckLogin).Methods("GET")
	return r
}
