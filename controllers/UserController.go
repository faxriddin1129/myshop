package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"myshop/models"
	"myshop/repository"
	"myshop/utils"
	"net/http"
	"strconv"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
		userModel, _ := models.GetUserByPhone(bodyParams.Phone)
		err := bcrypt.CompareHashAndPassword([]byte(userModel.PasswordHash), []byte(bodyParams.Password))
		if err != nil {
			utils.RespondWithError(w, http.StatusForbidden, map[string]string{"msg": "The password or phone number is incorrect"})
		} else {

			token := utils.GenerateToken(strconv.Itoa(int(int64(userModel.ID))))

			currentTime := time.Now()
			expireTime := currentTime.Add(10 * 24 * time.Hour)

			IPv6 := r.RemoteAddr
			Device := r.UserAgent()

			tokenModel := models.Token{}
			tokenModel.UserId = int64(userModel.ID)
			tokenModel.Token = token
			tokenModel.Expire = expireTime
			tokenModel.Ip = IPv6
			tokenModel.Device = Device
			tokenModel.CreateAccessToken()

			utils.RespondWithSuccess(w, map[string]string{
				"msg":    "Success",
				"Token":  token,
				"Expire": expireTime.Format("2006-01-02 15:04:05"),
				"IPv6":   IPv6,
				"Device": Device,
			}, nil)
		}
	}
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	userModel := utils.Auth(r.Context())
	utils.RespondWithSuccess(w, nil, userModel)
}
