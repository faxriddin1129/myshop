package services

import "net/http"

var httpClient = &http.Client{}

func Get(url string) interface{} {

	data := make(map[string]string)
	data["url"] = url
	return data

}

func Post() {

}

func Put() {

}

func Delete() {

}
