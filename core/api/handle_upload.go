package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// HandleUpload concerns request which upload either:
// - an image with the dimensions of 16x16 and are .png
// - a gif with the dimensions of 16x16
// if either the dimensions don't match or the file extensions
// get violated the endpoint returns a http.StatusBadRequest
func (a *Api) HandleUpload(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	format, ok := vars["formate"]
	if !ok {
		http.Error(w, "missing asset format", http.StatusBadRequest)
		return
	}
	// set may upload size to 1024 byte (16x16 image with rgba)
	r.ParseMultipartForm(maxUploadSize)

	file, _, err := r.FormFile("asset")
	if err != nil {
		logrus.Errorf("[HandleUpload] could not retrieve asset from request: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := a.assets.Store(format, file); err != nil {
		logrus.Errorf("[HandleUpload] could not store asset: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
