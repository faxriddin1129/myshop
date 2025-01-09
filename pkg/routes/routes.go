package routes

import (
	"MYSHOP/pkg/controllers"
	"github.com/gorilla/mux"
)

var UserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/user/get-me", controllers.GetMe).Methods("GET")
}
