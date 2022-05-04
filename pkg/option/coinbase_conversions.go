package option

// * This is a generated file, do not edit

// CoinbaseConversions are options for API requests.
type CoinbaseConversions struct {
	Amount    float64 `json:"amount" bson:"amount"`
	From      string  `json:"from" bson:"from"`
	Nonce     *string `json:"nonce" bson:"nonce"`
	ProfileId *string `json:"profile_id" bson:"profile_id"`
	To        string  `json:"to" bson:"to"`
}

// SetProfileId sets the ProfileId field on CoinbaseConversions.
func (opts *CoinbaseConversions) SetProfileId(profileId string) *CoinbaseConversions {
	opts.ProfileId = &profileId
	return opts
}

// SetFrom sets the From field on CoinbaseConversions.
func (opts *CoinbaseConversions) SetFrom(from string) *CoinbaseConversions {
	opts.From = from
	return opts
}

// SetTo sets the To field on CoinbaseConversions.
func (opts *CoinbaseConversions) SetTo(to string) *CoinbaseConversions {
	opts.To = to
	return opts
}

// SetAmount sets the Amount field on CoinbaseConversions.
func (opts *CoinbaseConversions) SetAmount(amount float64) *CoinbaseConversions {
	opts.Amount = amount
	return opts
}

// SetNonce sets the Nonce field on CoinbaseConversions.
func (opts *CoinbaseConversions) SetNonce(nonce string) *CoinbaseConversions {
	opts.Nonce = &nonce
	return opts
}
