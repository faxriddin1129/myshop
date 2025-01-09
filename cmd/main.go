package main

import (
	"MYSHOP/pkg/routes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"msg": "Welcome to My Shop!"}
	res, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
