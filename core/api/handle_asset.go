package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
