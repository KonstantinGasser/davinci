package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (a *Api) HandleUpdates(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	assestID, ok := vars["asset"]
	if !ok {
		http.Error(w, "missing asset-id", http.StatusBadRequest)
		return
	}
	formate, ok := vars["formate"]
	if !ok {
		http.Error(w, "missing format-type", http.StatusBadRequest)
		return
	}

	switch formate {
	case "img":
		img, err := a.assets.Image(assestID)
		if err != nil {
			logrus.Errorf("[HandleUpdates] could not load asset: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := a.matrixSvc.Print(img); err != nil {
			logrus.Errorf("[HandleUpdates] could not display image: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "gif":
		gif, err := a.assets.GIF(assestID)
		if err != nil {
			logrus.Errorf("[HandleUpdates] could not load asset: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := a.matrixSvc.Animate(gif); err != nil {
			logrus.Errorf("[HandleUpdates] could not run GIF: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "unknown file formate", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
