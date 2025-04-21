package router

import (
	"log"
	"main/data"
	"net/http"
)

func addJsonContentType(header http.Header) {
	header.Add(http.CanonicalHeaderKey("content-type"), "application/json")
}

func Init(db data.Datas, router *http.ServeMux) {

	if db == nil {
		log.Fatal("JSON DATABASE Failed to produce")
	}

	// get everything of the data
	router.HandleFunc("GET /", GetAll)
	// same as get
	router.HandleFunc("GET /products", GetAll)
	// get a product by id
	router.HandleFunc("GET /product/{id}", db.ByID)
	// echo
	router.HandleFunc("POST /echo", Echo)
	// adding list
	router.HandleFunc("POST /products", CreateNew())
	// get specific
	router.HandleFunc("GET /product/single", GetSpecific())
	// edit idea
	router.HandleFunc("PUT /product/{id}", PatchUpdateProduct())
	// delete  by id
	router.HandleFunc("DELETE /product/{id}", DeleteProduct())
}
