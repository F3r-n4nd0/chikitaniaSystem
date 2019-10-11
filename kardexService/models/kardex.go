package models

type Kardex struct {
	KardexId string `json:"kardex_id"`
	ProductId    string `json:"product_id"`
	CreateOn int64 `json:"create_on"`
	ProductPrice float32 `json:"product_price"`
}