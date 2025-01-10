package controllers

import (
	"MYSHOP/pkg/models"
	"MYSHOP/pkg/repository"
	"MYSHOP/pkg/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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
		utils.RespondWithError(w, http.StatusForbidden, errorMessage)
	} else {
		response, _ := models.GetUserByPhone(bodyParams.Phone)
		err := bcrypt.CompareHashAndPassword([]byte(response.PasswordHash), []byte(bodyParams.Password))
		if err != nil {
			utils.RespondWithError(w, http.StatusForbidden, map[string]string{"msg": "The password or phone number is incorrect"})
		} else {

			utils.RespondWithSuccess(w, map[string]string{
				"msg":    "Success",
				"Token":  "1",
				"Expire": "1",
			})
		}
	}
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithSuccess(w, map[string]string{"msg": "GetMe Success"})
}
