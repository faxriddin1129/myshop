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
