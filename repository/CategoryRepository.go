package repository

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"myshop/models"
	"myshop/utils"
	"net/http"
)

type CategoryValidateRepository struct {
	NameUz         string `json:"NameUz" validate:"required,max=255"`
	NameRu         string `json:"NameRu" validate:"required,max=255"`
	Type           string `json:"Type" validate:"required,max=255"`
	ImageUrl       string `json:"ImageUrl" validate:"max=255"`
	ParentID       *int64 `json:"ParentID"`
	FileId         int64  `json:"FileId"`
	CurrentFileUrl string `json:"CurrentFileUrl" validate:"max=255"`
}

func CategorySave(w http.ResponseWriter, r *http.Request) {
	bodyParams := &CategoryValidateRepository{}
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

	category := models.Category{}
	category.NameUz = bodyParams.NameUz
	category.NameRu = bodyParams.NameRu
	category.Type = bodyParams.Type
	category.ParentID = bodyParams.ParentID
	category.FileId = bodyParams.FileId
	category.CurrentFileUrl = currentUrl

	model := models.CreateCategory(&category)
	utils.RespondWithSuccess(w, nil, model)
}

func CategoryUpdate(w http.ResponseWriter, r *http.Request, modelCategory *models.Category) {

	bodyParams := &CategoryValidateRepository{}
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

	modelCategory.NameUz = bodyParams.NameUz
	modelCategory.NameRu = bodyParams.NameRu
	modelCategory.Type = bodyParams.Type
	modelCategory.ParentID = bodyParams.ParentID
	modelCategory.FileId = bodyParams.FileId

	currentUrl := ""
	if bodyParams.FileId != 0 {
		fileModel, db := models.GetFileById(bodyParams.FileId)
		if db.RowsAffected == 0 {
			utils.RespondWithError(w, http.StatusNotFound, map[string]string{"message": "File not found"})
			return
		}
		currentUrl = fileModel.CurrentUrl
	}
	modelCategory.CurrentFileUrl = currentUrl

	_, err = models.CategoryUpdate(*modelCategory)
	if err != nil {
		return
	}

	utils.RespondWithSuccess(w, nil, modelCategory)
}

func CategoryDelete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "CategoryDelete!", http.StatusNotImplemented)
}
