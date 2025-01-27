package controllers

import (
	"myshop/models"
	"myshop/repository"
	"myshop/utils"
	"net/http"
)

func ProductGetAll(w http.ResponseWriter, r *http.Request) {
	res := models.ProductGetAll()
	utils.RespondWithSuccess(w, nil, res)
	return
}

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	repository.ProductSave(w, r)
	return
}
