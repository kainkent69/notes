package router

import (
	"encoding/json"
	"io"
	"log"
	"main/data"
	"net/http"
	"strings"
)

func ServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Header().Add(http.CanonicalHeaderKey("content-type"), "application/json")
	w.Write([]byte("{ error: true, type: SERVER_ERR}"))
}
func Echo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	log.Println("Token Data")
	if err != nil {
		ServerError(w, r)
		return
	}
	// echo the data back
	w.Header().Add(http.CanonicalHeaderKey("content-type"), "application/json")
	w.Write(body)
	err = r.Body.Close()
	if err != nil {
		ServerError(w, r)
		return
	}
}

func InvalidInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
	w.Header().Add(http.CanonicalHeaderKey("content-type"), "application/json")
	w.Write([]byte("{ error: Invalid Data}"))
	r.Body.Close()
}

// create a new data
func CreateNew(datas data.Datas) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		product := data.Product{}
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Printf("Not decoded %v", err)
			ServerError(w, r)
			return
		}
		// validate the data
		if len([]byte(product.Info.Name)) <= 0 {
			log.Print("Name Is Invalid")
			InvalidInfo(w, r)
			return
		} else if product.Info.Price <= 0.0 {
			InvalidInfo(w, r)

			log.Printf("Price is Invalid: %v", product.Info.Price)
			return
		} else if strings.Compare(product.Info.Type, "") == 0 {

			log.Print("Type Is Invalid")
			InvalidInfo(w, r)
			return
		} else if product.Tags == nil || len(product.Tags) <= 0 {
			// the meta
			product.Tags = []string{}
		}
		product.Id = len(datas)
		datas.Append(product)
		w.WriteHeader(201)
		r.Body.Close()
		err = datas.Save()
		if err != nil {
			log.Printf("Not saved %v", err)
			ServerError(w, r)
		}
	}
}
