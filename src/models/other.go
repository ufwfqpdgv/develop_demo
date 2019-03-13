package models

type VipLevelCode int

const (
	VipLevelCode_Gold VipLevelCode = iota + 1
	VipLevelCode_Diamond
)

type DataTemp1 struct {
	ProductItem *Product `json:"product_item,omitempty"`
}

type Product struct {
	ProductId int64  `json:"product_id,omitempty"`
	Price     int    `json:"price,omitempty"`
	Detail    string `json:"detail,omitempty"`
	Extra     string `json:"extra,omitempty"`
}
