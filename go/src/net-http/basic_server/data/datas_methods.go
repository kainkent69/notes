package data

import (
	"encoding/json"
	"errors"
	"iter"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// make an each iterator
func (d *Datas) IterEach() iter.Seq2[int, *Product] {
	return func(yield func(int, *Product) bool) {
		for i := range *d {
			if !yield(i, &((*d)[i])) {
				return
			}
		}
	}
}

// make looping each and return either *Product or  Datas
func (d *Datas) Each(loop func(p *Product) bool) Datas {
	data := make([]Product, 0)
	for _, v := range d.IterEach() {
		if loop(v) {
			data = append(data, *v)
		}

	}

	return data
}

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

// adding new data to the Product slice/array
func (d *Datas) Append(p Product) {
	*d = append(*d, p)
}

// save the current changes to the file database
func (d *Datas) Save() error {
	path := "./data/data.json"
	mut := new(sync.Mutex)
	mut.Lock()
	defer mut.Unlock()
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

// product
func (d *Datas) Edit(id int, newproduct Product) (Product, error) {
	old := (*d)[id]
	product := &(*d)[id]
	sample := d.Each(func(p *Product) bool {
		if p.Id == id {
			return true
		}
		return false
	})

	// update the data
	if len(sample) <= 0 {
		return Product{}, errors.New("data.Datas.Edit: Nothing To Edit")
	} else {
		if len(newproduct.Tags) > 0 {
			// product tags
			product.Tags = newproduct.Tags
			log.Printf("Tags Updated ")
		}
		// product name
		if len([]byte(newproduct.Info.Name)) > 0 {

			product.Info.Name = newproduct.Info.Name
			log.Printf("Name Updated %v", newproduct.Info.Name)
		}
		// product type
		if len([]byte(newproduct.Info.Type)) > 0 {
			product.Info.Type = newproduct.Info.Type
			log.Printf("Type Update %v", newproduct.Info.Type)
		}
		// product  price
		if newproduct.Info.Price > 0.0 {
			product.Info.Price = newproduct.Info.Price
			log.Printf("Pricek Update %v", newproduct.Info.Price)

		}
	}

	(*d)[id] = *product

	log.Printf("\n\n\n\nThe Updated \n\n%#+v\n\n\n\n", *d)
	err := d.Save()
	if err != nil {
		return Product{}, err
	}
	return old, nil

}
