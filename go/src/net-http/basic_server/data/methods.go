package data

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func (d *Datas) ByID(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add(http.CanonicalHeaderKey("content-type"), "application/json")
	// range over the slices
	pathID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
	}

	for _, s := range *d {
		if s.Id == pathID {
			result, err := json.Marshal(Product(s))
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

func NewProduct(Id int, info ProductInfo, meta ProductMeta, tags []string, others []string) *Product {
	data := &Product{Id: Id, Tags: tags, Info: info, Meta: meta, Others: others}
	return data
}

func (d *Datas) Append(p Product) {
	*d = append(*d, p)
}

func (d *Datas) Save() error {
	path := "./data/data.json"

	// parse the data
	datastr, err := json.MarshalIndent(*d, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, datastr, 7)
	if err != nil {
		return err
	}
	// return safetly
	return nil
}
