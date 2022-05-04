package option

// * This is a generated file, do not edit

// CoinbaseAccountDeposit are options for API requests.
type CoinbaseAccountDeposit struct {
	Amount            float64 `json:"amount" bson:"amount"`
	CoinbaseAccountId string  `json:"coinbase_account_id" bson:"coinbase_account_id"`
	Currency          string  `json:"currency" bson:"currency"`
	ProfileId         *string `json:"profile_id" bson:"profile_id"`
}

// SetProfileId sets the ProfileId field on CoinbaseAccountDeposit.
func (opts *CoinbaseAccountDeposit) SetProfileId(profileId string) *CoinbaseAccountDeposit {
	opts.ProfileId = &profileId
	return opts
}

// SetAmount sets the Amount field on CoinbaseAccountDeposit.
func (opts *CoinbaseAccountDeposit) SetAmount(amount float64) *CoinbaseAccountDeposit {
	opts.Amount = amount
	return opts
}

// SetCoinbaseAccountId sets the CoinbaseAccountId field on CoinbaseAccountDeposit.
func (opts *CoinbaseAccountDeposit) SetCoinbaseAccountId(coinbaseAccountId string) *CoinbaseAccountDeposit {
	opts.CoinbaseAccountId = coinbaseAccountId
	return opts
}

// SetCurrency sets the Currency field on CoinbaseAccountDeposit.
func (opts *CoinbaseAccountDeposit) SetCurrency(currency string) *CoinbaseAccountDeposit {
	opts.Currency = currency
	return opts
}
