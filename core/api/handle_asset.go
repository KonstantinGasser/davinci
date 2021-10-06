package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (a *Api) HandleAsset(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	assetID, ok := vars["asset"]
	if !ok {
		http.Error(w, "missing asset ID", http.StatusBadRequest)
		return
	}

	file, format, err := a.assets.Load(assetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", fmt.Sprintf("image/%s", format))
	w.Write(file)
}

func (a *Api) HandleAssetList(w http.ResponseWriter, r *http.Request) {

	assets, err := a.assets.List()
	if err != nil {
		logrus.Errorf("[HandleAssetList] could not list asset dir: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(assets); err != nil {
		logrus.Errorf("[HandleAssetList] could not encode to json: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
