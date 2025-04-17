package models

import "time"

type StyleColor string
type MerchGroup string

type SizeSkuMap struct {
	SkuId string `json:"skuId,omitempty"`
	Gtin  string `json:"gtin,omitempty"`
}

type NikeSku struct {
	Id                 int64                 `json:"id,omitempty"`
	StyleColor         string                `json:"styleColor"`
	MerchGroup         string                `json:"merchGroup"`
	ProductName        string                `json:"productName,omitempty"`
	ProductDescription string                `json:"productDescription,omitempty"`
	ProductId          string                `json:"productId,omitempty"`
	LaunchId           string                `json:"launchId,omitempty"`
	Price              string                `json:"price,omitempty"`
	CurrentPrice       string                `json:"currentPrice,omitempty"`
	PublishType        string                `json:"publishType,omitempty"`
	CommerceStartTime  time.Time             `json:"commerceStartTime,omitempty"`
	CountryExclusion   []string              `json:"countryExclusion,omitempty"`
	SizeSkuMap         map[string]SizeSkuMap `json:"sizeSkuMap,omitempty"`
	QuantityLimit      int                   `json:"quantityLimit,omitempty"`
	Status             string                `json:"status,omitempty"`
	Exclusive          bool                  `json:"exclusive,omitempty"`
	Discountability    bool                  `json:"discountability,omitempty"`
}

type TempSku struct {
	Id         int64  `json:"id,omitempty"`
	StyleColor string `json:"styleColor"`
	MerchGroup string `json:"merchGroup"`
	Action     string `json:"action"`
}
