package data

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type ProductInfo struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type ProductMeta struct {
	Code  string  `json:"code"`
	Stock float64 `json:"stock"`
}
type Product struct {
	Id     int      `json:"id"`
	Tags   []string `json:"tags"`
	Others []string `json:"-"`

	// the other part
	Info ProductInfo `json:"info"`
	Meta ProductMeta `json:"meta"`
}

type Datas []Product

func ReadData() []byte {
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
	dbstr := ReadData()
	db := make(Datas, 0)
	if dbstr == nil {
		return nil
	}
	err := json.Unmarshal(dbstr, &db)
	if err != nil {
		return nil
	}
	return db
}
