package controllers

import (
	"myshop/models"
	"myshop/utils"
	"net/http"
)

func ProductGetAll(w http.ResponseWriter, r *http.Request) {
	res := models.ProductGetAll()
	utils.RespondWithSuccess(w, nil, res)
}
