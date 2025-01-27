package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"myshop/routes"
	"myshop/utils"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcome).Methods("GET")
	routes.ValidRoutes(router)
	http.Handle("/", router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	fmt.Println("Listening on :8001")
	log.Fatal(http.ListenAndServe(":8001", handler))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithSuccess(w, map[string]string{
		"msg": "Welcome to My Shop!",
	}, nil)
}

func init() {
	location, _ := time.LoadLocation("Asia/Tashkent")
	time.Local = location
}
