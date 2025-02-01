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

	router.Handle("/brand", middleware.AuthMiddleware(http.HandlerFunc(controllers.BrandCreate))).Methods("POST")
	router.Handle("/brand", middleware.AuthMiddleware(http.HandlerFunc(controllers.BrandGetAll))).Methods("GET")
	router.Handle("/brand", middleware.AuthMiddleware(http.HandlerFunc(controllers.BrandUpdate))).Methods("PUT")
	router.Handle("/brand", middleware.AuthMiddleware(http.HandlerFunc(controllers.BrandDelete))).Methods("DELETE")

	router.Handle("/product", middleware.AuthMiddleware(http.HandlerFunc(controllers.ProductCreate))).Methods("POST")
	router.Handle("/product", middleware.AuthMiddleware(http.HandlerFunc(controllers.ProductGetAll))).Methods("GET")
	router.Handle("/product", middleware.AuthMiddleware(http.HandlerFunc(controllers.ProductUpdate))).Methods("PUT")
	router.Handle("/product-view", middleware.AuthMiddleware(http.HandlerFunc(controllers.ProductView))).Methods("GET")
	router.Handle("/product-add-image", middleware.AuthMiddleware(http.HandlerFunc(controllers.ProductAddImage))).Methods("POST")
	router.Handle("/product-remove-image", middleware.AuthMiddleware(http.HandlerFunc(controllers.ProductImageDelete))).Methods("DELETE")

	router.Handle("/banner", middleware.AuthMiddleware(http.HandlerFunc(controllers.BannerCreate))).Methods("POST")
	router.Handle("/banner", middleware.AuthMiddleware(http.HandlerFunc(controllers.BannerGetAll))).Methods("GET")
	router.Handle("/banner-view", middleware.AuthMiddleware(http.HandlerFunc(controllers.BannerGetById))).Methods("GET")
	router.Handle("/banner", middleware.AuthMiddleware(http.HandlerFunc(controllers.BannerDelete))).Methods("DELETE")

	router.Handle("/config", middleware.AuthMiddleware(http.HandlerFunc(controllers.ConfigUpdate))).Methods("PUT")
	router.Handle("/config", middleware.AuthMiddleware(http.HandlerFunc(controllers.ConfigGet))).Methods("GET")

}
