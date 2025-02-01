package repository

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"myshop/models"
	"myshop/utils"
	"net/http"
)

type BrandValidateRepository struct {
	NameUz         string `json:"NameUz" validate:"required,max=255"`
	NameRu         string `json:"NameRu" validate:"required,max=255"`
	FileId         int64  `json:"FileId"`
	CurrentFileUrl string `json:"CurrentFileUrl" validate:"max=255"`
}

func BrandCreate(w http.ResponseWriter, r *http.Request) {
	bodyParams := &BrandValidateRepository{}
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

	brand := models.Brand{}
	brand.NameUz = bodyParams.NameUz
	brand.NameRu = bodyParams.NameRu
	brand.FileId = bodyParams.FileId
	brand.CurrentFileUrl = currentUrl

	model := models.BrandCreate(&brand)
	utils.RespondWithSuccess(w, nil, model)
}

func BrandUpdate(w http.ResponseWriter, r *http.Request, brandCategory *models.Brand) {

	bodyParams := &BrandValidateRepository{}
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

	brandCategory.NameUz = bodyParams.NameUz
	brandCategory.NameRu = bodyParams.NameRu
	brandCategory.FileId = bodyParams.FileId

	currentUrl := ""
	if bodyParams.FileId != 0 {
		fileModel, db := models.GetFileById(bodyParams.FileId)
		if db.RowsAffected == 0 {
			utils.RespondWithError(w, http.StatusNotFound, map[string]string{"message": "File not found"})
			return
		}
		currentUrl = fileModel.CurrentUrl
	}
	brandCategory.CurrentFileUrl = currentUrl
	_, err = models.BrandUpdate(*brandCategory)

	utils.RespondWithSuccess(w, nil, brandCategory)
}
