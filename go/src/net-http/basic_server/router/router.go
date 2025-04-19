package router

import (
	"log"
	"main/data"
	"main/middlewares"
	"net/http"
)

func addJsonContentType(header http.Header) {
	header.Add(http.CanonicalHeaderKey("content-type"), "application/json")
}

func Init(db data.Datas, router *http.ServeMux) {

	if db == nil {
		log.Fatal("JSON DATABASE Failed to produce")
	}
	server := http.Server{
		Addr:    ":8000",
		Handler: middlewares.Log(router),
	}

	// get everything of the data
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		addJsonContentType(header)
		w.Write(data.ReadData())
		r.Body.Close()
	})

	// same as get /
	router.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		addJsonContentType(header)
		w.Write(data.ReadData())
		r.Body.Close()
	})

	router.HandleFunc("GET /product/{id}", db.ByID)
	log.Printf("Server is running at port 8000")
	log.Fatal(server.ListenAndServe())

}
