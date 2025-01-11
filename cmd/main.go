package main

import (
	"MYSHOP/pkg/routes"
	"MYSHOP/pkg/utils"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcome)
	routes.UserRoutes(router)
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
