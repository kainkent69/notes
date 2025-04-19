package main

import (
	"bufio"
	"encoding/json"
	"log"
	"main/data"
	"main/middlewares"
	"net/http"
	"os"
	"strconv"
)

var router http.ServeMux

type Datas []data.Product

func getData() []byte {
	path := "./data/data.json"
	fs, err := os.Open(path)
	if err != nil {
		log.Fatalf("errored: %v", err)
	}
	scanner := bufio.NewScanner(fs)
	buf := make([]byte, 0)

	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			return nil
		}
		buf = append(buf, []byte(scanner.Text())...)
	}
	return buf
}

func DB() Datas {
	dbstr := getData()
	db := make([]data.Product, 0)
	if dbstr == nil {
		return nil
	}
	err := json.Unmarshal(dbstr, &db)
	if err != nil {
		return nil
	}
	return db
}

func addJsonContentType(header http.Header) {
	header.Add(http.CanonicalHeaderKey("content-type"), "application/json")
}

func (d *Datas) ByID(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	addJsonContentType(header)
	// range over the slices
	pathID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
	}

	for _, s := range *d {
		if s.Id == pathID {
			result, err := json.Marshal(data.Product(s))
			if err != nil {
				w.WriteHeader(500)
				r.Body.Close()
				return
			} else {
				w.WriteHeader(202)
				w.Write(result)
				r.Body.Close()
				return
			}

		}
	}

	w.WriteHeader(404)
	w.Write([]byte("NOT FOUND"))
	r.Body.Close()
}
func main() {
	db := DB()
	if db == nil {
		log.Fatal("JSON DATABASE Failed to produce")
	}
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8000",
		Handler: middlewares.Log(router),
	}

	// get everything of the data
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		addJsonContentType(header)
		w.Write(getData())
		r.Body.Close()
	})

	// same as get /
	router.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		addJsonContentType(header)
		w.Write(getData())
		r.Body.Close()
	})

	router.HandleFunc("GET /product/{id}", db.ByID)
	log.Printf("Server is running at port 8000")
	log.Fatal(server.ListenAndServe())

}
