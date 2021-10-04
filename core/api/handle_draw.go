package api

import (
	"encoding/json"
	"net/http"

	"github.com/KonstantinGasser/davinci/core/pkg/asset"
)

func (a *Api) HandleDraw(w http.ResponseWriter, r *http.Request) {
	var reqB [][]struct {
		I, J int
		Hex  string
	}
	if err := json.NewDecoder(r.Body).Decode(&reqB); err != nil {
		http.Error(w, "could not decode", http.StatusBadRequest)
		return
	}

	img := asset.From2dArray(reqB)

	if err := a.matrixSvc.Print(img); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
