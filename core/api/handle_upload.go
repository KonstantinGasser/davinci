package api

import "net/http"

// HandleUpload concerns request which upload either:
// - an image with the dimensions of 16x16 and are .png
// - a gif with the dimensions of 16x16
// if either the dimensions don't match or the file extensions
// get violated the endpoint returns a http.StatusBadRequest
func (a *Api) HandleUpload(w http.ResponseWriter, r *http.Request) {

}
