package data

type ProductInfo struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"prince"`
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
