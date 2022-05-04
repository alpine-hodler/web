package option

// * This is a generated file, do not edit

// CoinbaseAccountTransfers are options for API requests.
type CoinbaseAccountTransfers struct {
	After  *string `json:"after" bson:"after"`
	Before *string `json:"before" bson:"before"`
	Limit  *int    `json:"limit" bson:"limit"`
	Type   *string `json:"type" bson:"type"`
}

// SetBefore sets the Before field on CoinbaseAccountTransfers.
func (opts *CoinbaseAccountTransfers) SetBefore(before string) *CoinbaseAccountTransfers {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on CoinbaseAccountTransfers.
func (opts *CoinbaseAccountTransfers) SetAfter(after string) *CoinbaseAccountTransfers {
	opts.After = &after
	return opts
}

// SetLimit sets the Limit field on CoinbaseAccountTransfers.
func (opts *CoinbaseAccountTransfers) SetLimit(limit int) *CoinbaseAccountTransfers {
	opts.Limit = &limit
	return opts
}

// SetType sets the Type field on CoinbaseAccountTransfers.
func (opts *CoinbaseAccountTransfers) SetType(typ string) *CoinbaseAccountTransfers {
	opts.Type = &typ
	return opts
}
