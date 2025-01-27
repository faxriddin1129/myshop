package controllers

import (
	"myshop/models"
	"myshop/repository"
	"myshop/utils"
	"net/http"
	"strconv"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	repository.CategoryCreate(w, r)
	return
}

func CategoryGetAll(w http.ResponseWriter, r *http.Request) {
	ID := 0
	ParentId := r.URL.Query().Get("id")
	if ParentId != "" {
		ID, _ = strconv.Atoi(ParentId)
	}
	res := models.CategoryGetAll(ID)
	utils.RespondWithSuccess(w, nil, res)
	return
}

func CategoryUpdate(w http.ResponseWriter, r *http.Request) {

	ID := 0
	QueryId := r.URL.Query().Get("id")
	if QueryId != "" {
		ID, _ = strconv.Atoi(QueryId)
	} else {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, map[string]string{"msg": "Id is required"})
		return
	}

	model, _ := models.CategoryGetById(int64(ID))
	if model.ID == 0 {
		utils.RespondWithError(w, http.StatusNotFound, map[string]string{"msg": "Category not found"})
		return
	}

	repository.CategoryUpdate(w, r, model)
	return
}

func CategoryDelete(w http.ResponseWriter, r *http.Request) {

	ID := 0
	QueryId := r.URL.Query().Get("id")
	if QueryId != "" {
		ID, _ = strconv.Atoi(QueryId)
	} else {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, map[string]string{"msg": "Id is required"})
		return
	}

	model, _ := models.CategoryGetById(int64(ID))
	if model.ID == 0 {
		utils.RespondWithError(w, http.StatusNotFound, map[string]string{"msg": "Category not found"})
		return
	}

	models.CategoryDelete(int64(ID))
	utils.RespondWithSuccess(w, map[string]string{"msg": "Deleted category"}, nil)
	return
}
