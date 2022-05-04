package option

// * This is a generated file, do not edit

// CoinbaseAccountWithdrawal are options for API requests.
type CoinbaseAccountWithdrawal struct {
	Amount            float64 `json:"amount" bson:"amount"`
	CoinbaseAccountId string  `json:"coinbase_account_id" bson:"coinbase_account_id"`
	Currency          string  `json:"currency" bson:"currency"`
	ProfileId         *string `json:"profile_id" bson:"profile_id"`
}

// SetProfileId sets the ProfileId field on CoinbaseAccountWithdrawal.
func (opts *CoinbaseAccountWithdrawal) SetProfileId(profileId string) *CoinbaseAccountWithdrawal {
	opts.ProfileId = &profileId
	return opts
}

// SetAmount sets the Amount field on CoinbaseAccountWithdrawal.
func (opts *CoinbaseAccountWithdrawal) SetAmount(amount float64) *CoinbaseAccountWithdrawal {
	opts.Amount = amount
	return opts
}

// SetCoinbaseAccountId sets the CoinbaseAccountId field on CoinbaseAccountWithdrawal.
func (opts *CoinbaseAccountWithdrawal) SetCoinbaseAccountId(coinbaseAccountId string) *CoinbaseAccountWithdrawal {
	opts.CoinbaseAccountId = coinbaseAccountId
	return opts
}

// SetCurrency sets the Currency field on CoinbaseAccountWithdrawal.
func (opts *CoinbaseAccountWithdrawal) SetCurrency(currency string) *CoinbaseAccountWithdrawal {
	opts.Currency = currency
	return opts
}
