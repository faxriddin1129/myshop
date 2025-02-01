package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"myshop/models"
	"myshop/utils"
	"net/http"
)

func ConfigUpdate(w http.ResponseWriter, r *http.Request) {
	bodyParams := &models.ConfigValidate{}
	utils.ParseBody(r, bodyParams)
	validate := validator.New()
	err := validate.Struct(bodyParams)
	if err != nil {
		errorMessage := map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage[err.Field()] = fmt.Sprintf("%s: %s", err.Field(), err.Tag())
		}
		utils.RespondWithError(w, http.StatusUnprocessableEntity, errorMessage)
		return
	}

	config := &models.Config{}
	config.ID = 1
	config.Phone = bodyParams.Phone
	config.Email = bodyParams.Email
	config.Telegram = bodyParams.Telegram
	config.TelegramNick = bodyParams.TelegramNick
	config.Instagram = bodyParams.Instagram
	config.Facebook = bodyParams.Facebook
	config.Youtube = bodyParams.Youtube
	config.Address = bodyParams.Address
	config.PublicOffer = bodyParams.PublicOffer
	config.FooterText = bodyParams.FooterText

	_, err = models.ConfigUpdate(*config)
	utils.RespondWithSuccess(w, nil, config)
	return
}

func ConfigGet(w http.ResponseWriter, r *http.Request) {
	res, _ := models.ConfigGetById(1)
	utils.RespondWithSuccess(w, nil, res)
	return
}
