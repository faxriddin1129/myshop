package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"myshop/models"
	"myshop/repository"
	"myshop/utils"
	"net/http"
	"strconv"
)

func ProductGetAll(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 10
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}
	offset := (page - 1) * limit
	res, total := models.ProductGetAll(limit, offset)

	categories := models.CategoryGetAllNoParents()
	categoryMap := make(map[uint]models.Category)
	for _, category := range categories {
		categoryMap[category.ID] = category
	}

	brands := models.BrandGetAll()
	brandMap := make(map[uint]models.Brand)
	for _, brand := range brands {
		brandMap[brand.ID] = brand
	}

	productRes := make([]repository.ProductResponse, len(res))
	for i, product := range res {
		productRes[i] = repository.ProductResponse{
			Product:  product,
			Category: categoryMap[uint(product.CategoryId)],
			Brand:    brandMap[uint(product.BrandId)],
		}
	}

	utils.RespondWithSuccess(w, nil, map[string]interface{}{
		"data":      productRes,
		"page":      page,
		"limit":     limit,
		"total":     total,
		"totalPage": (total + int64(limit) - 1) / int64(limit),
	})
}

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	repository.ProductSave(w, r)
	return
}

func ProductView(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, map[string]string{"msg": "id is empty"})
		return
	}

	productModel, _ := models.ProductGetById(int64(id))
	if productModel.ID == 0 {
		utils.RespondWithError(w, http.StatusInternalServerError, map[string]string{"msg": "Product not found"})
		return
	}

	categoryModel, _ := models.CategoryGetById(productModel.CategoryId)
	brandModel, _ := models.BrandGetById(productModel.BrandId)
	images, _ := models.GetProductImages(int64(productModel.ID))

	response := repository.ProductResponse{
		Product:  productModel,
		Category: categoryModel,
		Brand:    brandModel,
		Images:   images,
	}

	utils.RespondWithSuccess(w, nil, response)
	return
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, map[string]string{"msg": "id is empty"})
		return
	}
	productModel, _ := models.ProductGetById(int64(id))
	if productModel.ID == 0 {
		utils.RespondWithError(w, http.StatusInternalServerError, map[string]string{"msg": "Product not found"})
		return
	}

	repository.ProductUpdate(w, r, productModel)
	return
}

func ProductAddImage(w http.ResponseWriter, r *http.Request) {

	bodyParams := &repository.ProductImageValidate{}
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

	productModel, _ := models.ProductGetById(bodyParams.ProductId)
	if productModel.ID == 0 {
		utils.RespondWithError(w, http.StatusInternalServerError, map[string]string{"msg": "Product not found"})
		return
	}

	fileModel, _ := models.GetFileById(bodyParams.FileId)
	if fileModel.ID == 0 {
		utils.RespondWithError(w, http.StatusInternalServerError, map[string]string{"msg": "File not found"})
		return
	}

	model := models.ProductImage{}
	model.FileId = bodyParams.FileId
	model.ProductId = bodyParams.ProductId
	model.CurrentUrl = fileModel.CurrentUrl

	obj := models.CreateProductImages(&model)
	utils.RespondWithSuccess(w, nil, obj)
	return
}

func ProductImageDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, map[string]string{"msg": "id is empty"})
		return
	}

	productImageModel, _ := models.GetProductImageById(int64(id))
	if productImageModel.ID == 0 {
		utils.RespondWithError(w, http.StatusInternalServerError, map[string]string{"msg": "Image not found"})
		return
	}

	models.ProductImagesDelete(int64(id))
	utils.RespondWithSuccess(w, map[string]string{"msg": "Successfully deleted!"}, nil)
}
