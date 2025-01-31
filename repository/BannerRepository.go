package repository

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"myshop/models"
	"myshop/utils"
	"net/http"
)

type BannerValidateRepository struct {
	ProductId      int64  `json:"ProductId" validate:"required"`
	FileId         int64  `json:"FileId" validate:"required"`
	MobileFileId   int64  `json:"MobileFileId" validate:"required"`
	CurrentFileUrl string `json:"CurrentFileUrl" validate:"max=255"`
}

func BannerCreate(w http.ResponseWriter, r *http.Request) {
	bodyParams := &BannerValidateRepository{}
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

	currentUrl := ""
	if bodyParams.FileId != 0 {
		fileModel, db := models.GetFileById(bodyParams.FileId)
		if db.RowsAffected == 0 {
			utils.RespondWithError(w, http.StatusNotFound, map[string]string{"message": "File not found"})
			return
		}
		currentUrl = fileModel.CurrentUrl
	}

	currentUrlMobile := ""
	if bodyParams.MobileFileId != 0 {
		fileModel, db := models.GetFileById(bodyParams.MobileFileId)
		if db.RowsAffected == 0 {
			utils.RespondWithError(w, http.StatusNotFound, map[string]string{"message": "File not found"})
			return
		}
		currentUrlMobile = fileModel.CurrentUrl
	}

	banner := models.Banner{}
	banner.FileId = bodyParams.FileId
	banner.MobileFileId = bodyParams.MobileFileId
	banner.CurrentFileUrl = currentUrl
	banner.MobileCurrentFileUrl = currentUrlMobile
	banner.ProductId = bodyParams.ProductId

	model := models.BannerCreate(&banner)
	utils.RespondWithSuccess(w, nil, model)
}
