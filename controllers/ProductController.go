package controllers

import (
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
	utils.RespondWithSuccess(w, nil, map[string]interface{}{
		"data":      res,
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
