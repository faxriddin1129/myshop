package controllers

import (
	"myshop/models"
	"myshop/repository"
	"myshop/utils"
	"net/http"
	"strconv"
)

func BannerCreate(w http.ResponseWriter, r *http.Request) {
	repository.BannerCreate(w, r)
	return
}

func BannerGetAll(w http.ResponseWriter, r *http.Request) {
	res := models.BannerGetAll()
	utils.RespondWithSuccess(w, nil, res)
	return
}

func BannerGetById(w http.ResponseWriter, r *http.Request) {

	ID := 0
	QueryId := r.URL.Query().Get("id")
	if QueryId != "" {
		ID, _ = strconv.Atoi(QueryId)
	} else {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, map[string]string{"msg": "Id is required"})
		return
	}

	model, _ := models.BannerGetById(int64(ID))
	if model.ID == 0 {
		utils.RespondWithError(w, http.StatusNotFound, map[string]string{"msg": "Banner not found"})
		return
	}

	utils.RespondWithSuccess(w, nil, model)
	return
}

func BannerDelete(w http.ResponseWriter, r *http.Request) {

	ID := 0
	QueryId := r.URL.Query().Get("id")
	if QueryId != "" {
		ID, _ = strconv.Atoi(QueryId)
	} else {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, map[string]string{"msg": "Id is required"})
		return
	}

	model, _ := models.BannerGetById(int64(ID))
	if model.ID == 0 {
		utils.RespondWithError(w, http.StatusNotFound, map[string]string{"msg": "Banner not found"})
		return
	}

	models.BannerDelete(int64(ID))
	utils.RespondWithSuccess(w, map[string]string{"msg": "Deleted banner"}, nil)
	return
}
