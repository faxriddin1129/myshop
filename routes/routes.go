package routes

import (
	"github.com/gorilla/mux"
	"myshop/controllers"
	"myshop/middleware"
	"net/http"
)

var UserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.Handle("/user/get-me", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetMe))).Methods("GET")
}

var CategoryRoutes = func(router *mux.Router) {
	router.Handle("/category", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateCategory))).Methods("POST")
	router.Handle("/category", middleware.AuthMiddleware(http.HandlerFunc(controllers.CategoryGetAll))).Methods("GET")
	router.Handle("/category", middleware.AuthMiddleware(http.HandlerFunc(controllers.CategoryUpdate))).Methods("PUT")
}

var FileRoutes = func(router *mux.Router) {
	router.Handle("/file/upload", middleware.AuthMiddleware(http.HandlerFunc(controllers.FileUpload))).Methods("POST")
}
