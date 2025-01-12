package controllers

import (
	"MYSHOP/pkg/models"
	"MYSHOP/pkg/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func FileUpload(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Faylni yuklashda xatolik yuz berdi", http.StatusBadRequest)
		return
	}
	defer file.Close()

	allowedExtensions := []string{"jpg", "pdf", "png", "jpeg", "webp"}
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(header.Filename), "."))

	isAllowed := utils.InArray(allowedExtensions, ext)
	if !isAllowed {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, map[string]string{"msg": "This file format is not allowed to be uploaded"})
		return
	}

	currentTime := int(time.Now().Unix())
	timeString := strconv.Itoa(currentTime)

	storagePath, _ := utils.FileGetPath()
	filename := timeString + "." + ext
	filePath := filepath.Join(storagePath, filename)

	out, err := os.Create(filePath)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, map[string]string{"msg": "There was an error saving file"})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, map[string]string{"msg": "There was an error saving file"})
		return
	}

	fileModel := models.File{}
	fileModel.Name = filename
	fileModel.Path = "/" + storagePath
	fileModel.BaseUrl = utils.BASE_URL
	fileModel.CurrentUrl = utils.BASE_URL + "/" + filePath

	obj := models.CreateFileModel(&fileModel)

	utils.RespondWithSuccess(w, nil, obj)

}
