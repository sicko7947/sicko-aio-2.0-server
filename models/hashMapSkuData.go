package models

type HashMapSkuData struct {
	Id                 int64  `json:"id,omitempty"`
	StyleColor         string `json:"styleColor"`
	MerchGroup         string `json:"merchGroup"`
	ProductName        string `json:"productName,omitempty"`
	ProductDescription string `json:"productDescription,omitempty"`
	ProductId          string `json:"productId,omitempty"`
	LaunchId           string `json:"launchId,omitempty"`
	Price              string `json:"price,omitempty"`
	CurrentPrice       string `json:"currentPrice,omitempty"`
	PublishType        string `json:"publishType,omitempty"`
	CommerceStartTime  string `json:"commerceStartTime,omitempty"`
	CountryExclusion   string `json:"countryExclusion,omitempty"`
	SizeSkuMap         string `json:"sizeSkuMap,omitempty"`
	QuantityLimit      string `json:"quantityLimit,omitempty"`
	Status             string `json:"status,omitempty"`
	Exclusive          bool   `json:"exclusive,omitempty"`
	Discountability    bool   `json:"discountability,omitempty"`
}
