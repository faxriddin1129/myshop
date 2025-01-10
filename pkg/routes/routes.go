package routes

import (
	"MYSHOP/pkg/controllers"
	"MYSHOP/pkg/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

var UserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.Handle("/user/get-me", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetMe))).Methods("GET")
}
