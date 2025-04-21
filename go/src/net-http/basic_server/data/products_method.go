package data

import (
	"crypto/sha256"
	"errors"
	"math/rand"
	"strconv"
)

type ProductValidateError struct {
	EmptyName    error
	InvalidPrice error
	InvalidType  error
}

var (
	ProductValidateErr = ProductValidateError{
		EmptyName:    errors.New("data.Product: Product Name Empty"),
		InvalidPrice: errors.New("data.Product: Invalid Price"),
		InvalidType:  errors.New("data.Product: Invalid Type"),
	}
)

// Validate A new Product
func (p *Product) Validate() error {
	// validate the name
	switch {
	case len([]byte(p.Info.Name)) == 0:
		return ProductValidateErr.EmptyName

	case p.Info.Price <= 0 || p.Info.Price > 100:
		return ProductValidateErr.InvalidPrice

	case len([]byte(p.Info.Type)) == 0:
		return ProductValidateErr.InvalidType

	case len(p.Tags) == 0:
		p.Tags = []string{}
	default:
		goto exit
	}

exit:
	return nil
}

// Will Call Datas.Append
func (p *Product) ToDatas(datas Datas) {
	datas.Append(*p)
}

// Make A new Product
func NewProduct(Id int, info ProductInfo, meta ProductMeta, tags []string, others []string) *Product {
	data := &Product{Id: Id, Tags: tags, Info: info, Meta: meta, Others: others}
	h := sha256.New()
	h.Write([]byte(strconv.Itoa(rand.Intn(100000))))
	data.Meta.Code = string(h.Sum(nil))
	return data
}
