package repository

import (
	"MYSHOP/pkg/models"
	"MYSHOP/pkg/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CategoryValidateRepository struct {
	NameUz   string `json:"NameUz" validate:"required,max=255"`
	NameRu   string `json:"NameRu" validate:"required,max=255"`
	Type     string `json:"Type" validate:"required,max=255"`
	ImageUrl string `json:"ImageUrl" validate:"max=255"`
	ParentID *int64 `json:"ParentID"`
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

	category := models.Category{}
	category.NameUz = bodyParams.NameUz
	category.NameRu = bodyParams.NameRu
	category.Type = bodyParams.Type
	category.ParentID = bodyParams.ParentID

	model := models.CreateCategory(&category)
	utils.RespondWithSuccess(w, nil, model)
}
