package controllers

import (
	"MYSHOP/pkg/models"
	"MYSHOP/pkg/repository"
	"MYSHOP/pkg/utils"
	"net/http"
	"strconv"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	repository.CategorySave(w, r)
}

func CategoryGetAll(w http.ResponseWriter, r *http.Request) {

	ID := 0
	ParentId := r.URL.Query().Get("id")
	if ParentId != "" {
		ID, _ = strconv.Atoi(ParentId)
	}

	res := models.CategoryGetAll(ID)
	utils.RespondWithSuccess(w, nil, res)
}

func CategoryUpdate(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Implement me, Update", http.StatusNotImplemented)
}
