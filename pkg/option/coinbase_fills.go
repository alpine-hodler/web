package option

// * This is a generated file, do not edit

// CoinbaseFills are options for API requests.
type CoinbaseFills struct {
	After     *int    `json:"after" bson:"after"`
	Before    *int    `json:"before" bson:"before"`
	Limit     *int    `json:"limit" bson:"limit"`
	OrderId   *string `json:"order_id" bson:"order_id"`
	ProductId *string `json:"product_id" bson:"product_id"`
	ProfileId *string `json:"profile_id" bson:"profile_id"`
}

// SetOrderId sets the OrderId field on CoinbaseFills.
func (opts *CoinbaseFills) SetOrderId(orderId string) *CoinbaseFills {
	opts.OrderId = &orderId
	return opts
}

// SetProductId sets the ProductId field on CoinbaseFills.
func (opts *CoinbaseFills) SetProductId(productId string) *CoinbaseFills {
	opts.ProductId = &productId
	return opts
}

// SetProfileId sets the ProfileId field on CoinbaseFills.
func (opts *CoinbaseFills) SetProfileId(profileId string) *CoinbaseFills {
	opts.ProfileId = &profileId
	return opts
}

// SetLimit sets the Limit field on CoinbaseFills.
func (opts *CoinbaseFills) SetLimit(limit int) *CoinbaseFills {
	opts.Limit = &limit
	return opts
}

// SetBefore sets the Before field on CoinbaseFills.
func (opts *CoinbaseFills) SetBefore(before int) *CoinbaseFills {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on CoinbaseFills.
func (opts *CoinbaseFills) SetAfter(after int) *CoinbaseFills {
	opts.After = &after
	return opts
}
