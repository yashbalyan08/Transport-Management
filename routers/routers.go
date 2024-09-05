package router

import (
	"github.com/yashbalyan08/system/controllers"
	"github.com/yashbalyan08/system/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")
	r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")

	// Protected routes (require authentication)
	protected := r.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	// Add any protected routes here, e.g.:
	// protected.HandleFunc("/dashboard", controllers.DashboardHandler).Methods("GET")

	return r
}
