package router

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/data"
	"net/http"
	"strconv"
	"strings"
)

type InvalidInfoT struct {
	Name  string
	Cause error
	Msg   string
}

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

func InvalidInfo(w http.ResponseWriter, r *http.Request, info InvalidInfoT) {
	by, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		ServerError(w, r)
		return
	}
	w.WriteHeader(400)
	w.Header().Add(http.CanonicalHeaderKey("content-type"), "application/json")
	w.Write(by)
	r.Body.Close()
}

// create a new data
func CreateNew() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		datas := data.DB()
		info := InvalidInfoT{}
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
			InvalidInfo(w, r, info)
			return
		} else if product.Info.Price <= 0.0 {
			InvalidInfo(w, r, info)
			log.Printf("Price is Invalid: %v", product.Info.Price)
			return
		} else if strings.Compare(product.Info.Type, "") == 0 {

			log.Print("Type Is Invalid")
			InvalidInfo(w, r, info)
			return
		} else if product.Tags == nil || len(product.Tags) <= 0 {
			// the meta
			product.Tags = []string{}
		}

		// do meta

		{
			sub := make([]byte, 0)
			h := sha256.New()
			h.Write([]byte(strconv.Itoa(product.Id)))
			b := h.Sum(sub)
			product.Meta.Code = string(b)
		}

		product.Id = len(datas)
		datas.Append(product)
		w.WriteHeader(201)
		buf, err := json.Marshal(product)
		if err != nil {
			InvalidInfo(w, r, info)
			return
		}

		w.Write(fmt.Appendf(make([]byte, 0), "{%q: true, %q : %v}", "reacted", "new", string(buf)))
		err = datas.Save()
		r.Body.Close()
		log.Printf("\n\nDatas\n\n %v\n", datas)

		if err != nil {
			log.Printf("Not saved %v", err)
			ServerError(w, r)
		}
	}
}
