package router

import (
	"main/data"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	addJsonContentType(header)
	w.Write(data.ReadData())
	r.Body.Close()
}
