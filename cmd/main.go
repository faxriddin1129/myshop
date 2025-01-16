package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"myshop/routes"
	"myshop/utils"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcome)
	routes.ValidRoutes(router)
	http.Handle("/", router)
	fmt.Println("Listening on :8001")
	log.Fatal(http.ListenAndServe(":8001", router))
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
