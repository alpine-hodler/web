package option

// * This is a generated file, do not edit

// CoinbaseProducts are options for API requests.
type CoinbaseProducts struct {
	Type *string `json:"type" bson:"type"`
}

// SetType sets the Type field on CoinbaseProducts.
func (opts *CoinbaseProducts) SetType(typ string) *CoinbaseProducts {
	opts.Type = &typ
	return opts
}
