package controllers

import (
	"MYSHOP/pkg/models"
	"MYSHOP/pkg/repository"
	"MYSHOP/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bodyParams := &repository.UserLoginStructRepository{}
	utils.ParseBody(r, bodyParams)
	validate := validator.New()
	err := validate.Struct(bodyParams)
	if err != nil {
		errorMessage := map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage[err.Field()] = fmt.Sprintf("%s: %s", err.Field(), err.Tag())
		}
		res, _ := json.Marshal(errorMessage)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(res)
	} else {
		response := map[string]string{
			"Token":  bodyParams.Phone,
			"Expire": bodyParams.Password,
		}
		res, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetMe(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	//response := map[string]string{
	//	"msg": "GetMe Success",
	//}
	response := models.GetAllUsers()
	res, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
