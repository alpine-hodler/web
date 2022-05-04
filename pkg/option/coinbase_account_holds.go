package option

// * This is a generated file, do not edit

// CoinbaseAccountHolds are options for API requests.
type CoinbaseAccountHolds struct {
	After  *string `json:"after" bson:"after"`
	Before *string `json:"before" bson:"before"`
	Limit  *int    `json:"limit" bson:"limit"`
}

// SetBefore sets the Before field on CoinbaseAccountHolds.
func (opts *CoinbaseAccountHolds) SetBefore(before string) *CoinbaseAccountHolds {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on CoinbaseAccountHolds.
func (opts *CoinbaseAccountHolds) SetAfter(after string) *CoinbaseAccountHolds {
	opts.After = &after
	return opts
}

// SetLimit sets the Limit field on CoinbaseAccountHolds.
func (opts *CoinbaseAccountHolds) SetLimit(limit int) *CoinbaseAccountHolds {
	opts.Limit = &limit
	return opts
}
