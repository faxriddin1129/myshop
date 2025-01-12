package controllers

import "net/http"

func FileUpload(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Implement me, Upload File", http.StatusNotImplemented)
}
