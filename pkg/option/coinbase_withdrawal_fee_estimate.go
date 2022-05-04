package option

// * This is a generated file, do not edit

// CoinbaseWithdrawalFeeEstimate are options for API requests.
type CoinbaseWithdrawalFeeEstimate struct {
	CryptoAddress *string `json:"crypto_address" bson:"crypto_address"`
	Currency      *string `json:"currency" bson:"currency"`
}

// SetCurrency sets the Currency field on CoinbaseWithdrawalFeeEstimate.
func (opts *CoinbaseWithdrawalFeeEstimate) SetCurrency(currency string) *CoinbaseWithdrawalFeeEstimate {
	opts.Currency = &currency
	return opts
}

// SetCryptoAddress sets the CryptoAddress field on CoinbaseWithdrawalFeeEstimate.
func (opts *CoinbaseWithdrawalFeeEstimate) SetCryptoAddress(cryptoAddress string) *CoinbaseWithdrawalFeeEstimate {
	opts.CryptoAddress = &cryptoAddress
	return opts
}
