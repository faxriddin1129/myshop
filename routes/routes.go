package routes

import (
	"github.com/gorilla/mux"
	"myshop/controllers"
	"myshop/middleware"
	"net/http"
)

var ValidRoutes = func(router *mux.Router) {

	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.Handle("/user/get-me", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetMe))).Methods("GET")
	router.Handle("/file/upload", middleware.AuthMiddleware(http.HandlerFunc(controllers.FileUpload))).Methods("POST")

	router.Handle("/category", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateCategory))).Methods("POST")
	router.Handle("/category", middleware.AuthMiddleware(http.HandlerFunc(controllers.CategoryGetAll))).Methods("GET")
	router.Handle("/category", middleware.AuthMiddleware(http.HandlerFunc(controllers.CategoryUpdate))).Methods("PUT")
	router.Handle("/category", middleware.AuthMiddleware(http.HandlerFunc(controllers.CategoryDelete))).Methods("DELETE")

	router.Handle("/brand", middleware.AuthMiddleware(http.HandlerFunc(controllers.BrandCategory))).Methods("POST")
	router.Handle("/brand", middleware.AuthMiddleware(http.HandlerFunc(controllers.BrandGetAll))).Methods("GET")
	router.Handle("/brand", middleware.AuthMiddleware(http.HandlerFunc(controllers.BrandUpdate))).Methods("PUT")
	router.Handle("/brand", middleware.AuthMiddleware(http.HandlerFunc(controllers.BrandDelete))).Methods("DELETE")

	router.Handle("/product", middleware.AuthMiddleware(http.HandlerFunc(controllers.ProductGetAll))).Methods("GET")
	router.Handle("/product", middleware.AuthMiddleware(http.HandlerFunc(controllers.ProductCreate))).Methods("POST")
}
