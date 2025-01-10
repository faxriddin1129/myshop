package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}) {
	reqBody, _ := io.ReadAll(r.Body)
	err := r.Body.Close()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqBody, v)
	if err != nil {
		panic(err)
	}
}

func RespondWithError(w http.ResponseWriter, statusCode int, message map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(message)
	w.WriteHeader(statusCode)
	_, err := w.Write(res)
	if err != nil {
		panic(err)
	}
}

func RespondWithSuccess(w http.ResponseWriter, message map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(message)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		panic(err)
	}
}
