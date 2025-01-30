package repository

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"myshop/models"
	"myshop/utils"
	"net/http"
)

type ProductValidateRepository struct {
	NameUz             string  `json:"NameUz" validate:"required,max=255"`
	NameRu             string  `json:"NameRu" validate:"required,max=255"`
	FileId             int64   `json:"FileId" validate:"required"`
	CurrentFileUrl     string  `json:"CurrentFileUrl" validate:"omitempty,max=255"`
	ShortDescriptionUz string  `json:"ShortDescriptionUz" validate:"required,max=500"`
	ShortDescriptionRu string  `json:"ShortDescriptionRu" validate:"required,max=500"`
	DescriptionUz      string  `json:"DescriptionUz" validate:"omitempty,max=10000"`
	DescriptionRu      string  `json:"DescriptionRu" validate:"omitempty,max=10000"`
	Count              int     `json:"Count" validate:"required,gte=0"`
	Price              float64 `json:"Price" validate:"required,gte=0"`
	DiscountPrice      float64 `json:"DiscountPrice" validate:"gte=0"`
	Status             int8    `json:"Status" validate:"required,oneof=1 2"`
	CategoryId         int64   `json:"CategoryId" validate:"required"`
	BrandId            int64   `json:"BrandId" validate:"required"`
	ProductionTime     string  `json:"ProductionTime" validate:"omitempty"`
}

type ProductImageValidate struct {
	FileId    int64 `json:"FileId" validate:"required"`
	ProductId int64 `json:"ProductId" validate:"required"`
}

type ProductResponse struct {
	Product  interface{} `json:"Product"`
	Category interface{} `json:"Category"`
	Brand    interface{} `json:"Brand"`
	Images   interface{} `json:"Images"`
}

func ProductSave(w http.ResponseWriter, r *http.Request) {
	bodyParams := &ProductValidateRepository{}
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

	product := models.Product{}
	product.FileId = bodyParams.FileId
	product.CurrentFileUrl = currentUrl
	product.NameUz = bodyParams.NameUz
	product.NameRu = bodyParams.NameRu
	product.ShortDescriptionUz = bodyParams.ShortDescriptionUz
	product.ShortDescriptionRu = bodyParams.ShortDescriptionRu
	product.DescriptionUz = bodyParams.DescriptionUz
	product.DescriptionRu = bodyParams.DescriptionRu
	product.Count = bodyParams.Count
	product.Price = bodyParams.Price
	product.DiscountPrice = bodyParams.DiscountPrice
	product.Status = bodyParams.Status
	product.CategoryId = bodyParams.CategoryId
	product.BrandId = bodyParams.BrandId
	product.ProductionTime = bodyParams.ProductionTime

	model := models.ProductCreate(&product)
	utils.RespondWithSuccess(w, nil, model)
	return
}

func ProductUpdate(w http.ResponseWriter, r *http.Request, modelProduct *models.Product) {

	bodyParams := &ProductValidateRepository{}
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

	modelProduct.NameUz = bodyParams.NameUz
	modelProduct.NameRu = bodyParams.NameRu
	modelProduct.ShortDescriptionUz = bodyParams.ShortDescriptionUz
	modelProduct.ShortDescriptionRu = bodyParams.ShortDescriptionRu
	modelProduct.DescriptionUz = bodyParams.DescriptionUz
	modelProduct.DescriptionRu = bodyParams.DescriptionRu
	modelProduct.Count = bodyParams.Count
	modelProduct.Price = bodyParams.Price
	modelProduct.DiscountPrice = bodyParams.DiscountPrice
	modelProduct.Status = bodyParams.Status
	modelProduct.CategoryId = bodyParams.CategoryId
	modelProduct.BrandId = bodyParams.BrandId
	modelProduct.ProductionTime = bodyParams.ProductionTime

	currentUrl := ""
	if bodyParams.FileId != 0 {
		fileModel, db := models.GetFileById(bodyParams.FileId)
		if db.RowsAffected == 0 {
			utils.RespondWithError(w, http.StatusNotFound, map[string]string{"message": "File not found"})
			return
		}
		currentUrl = fileModel.CurrentUrl
	}
	modelProduct.CurrentFileUrl = currentUrl

	_, err = models.ProductUpdate(*modelProduct)
	if err != nil {
		return
	}

	utils.RespondWithSuccess(w, nil, modelProduct)
}
