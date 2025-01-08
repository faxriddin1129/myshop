package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"msg": "Welcome to My Shop!"}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcome)
	http.Handle("/", router)
	log.Println("Listening on :8001")
	log.Fatal(http.ListenAndServe(":8001", router))
}
