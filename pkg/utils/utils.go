package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"time"
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

func RespondWithSuccess(w http.ResponseWriter, message map[string]string, model interface{}) {
	w.Header().Set("Content-Type", "application/json")
	var res []byte
	var err error
	if model != nil {
		res, _ = json.Marshal(model)
	} else {
		res, _ = json.Marshal(message)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		panic(err)
	}
}

func GenerateToken(id string) string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	_, err = rand.Read(bytes)
	token := hex.EncodeToString(bytes)
	currentTime := time.Now().Format("20060102150405")
	accessToken := id + "_" + token + currentTime
	return accessToken
}
