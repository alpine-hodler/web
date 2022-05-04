package option

// * This is a generated file, do not edit

// CoinbasePaymentMethodDeposit are options for API requests.
type CoinbasePaymentMethodDeposit struct {
	Amount          float64 `json:"amount" bson:"amount"`
	Currency        string  `json:"currency" bson:"currency"`
	PaymentMethodId string  `json:"payment_method_id" bson:"payment_method_id"`
	ProfileId       *string `json:"profile_id" bson:"profile_id"`
}

// SetProfileId sets the ProfileId field on CoinbasePaymentMethodDeposit.
func (opts *CoinbasePaymentMethodDeposit) SetProfileId(profileId string) *CoinbasePaymentMethodDeposit {
	opts.ProfileId = &profileId
	return opts
}

// SetAmount sets the Amount field on CoinbasePaymentMethodDeposit.
func (opts *CoinbasePaymentMethodDeposit) SetAmount(amount float64) *CoinbasePaymentMethodDeposit {
	opts.Amount = amount
	return opts
}

// SetPaymentMethodId sets the PaymentMethodId field on CoinbasePaymentMethodDeposit.
func (opts *CoinbasePaymentMethodDeposit) SetPaymentMethodId(paymentMethodId string) *CoinbasePaymentMethodDeposit {
	opts.PaymentMethodId = paymentMethodId
	return opts
}

// SetCurrency sets the Currency field on CoinbasePaymentMethodDeposit.
func (opts *CoinbasePaymentMethodDeposit) SetCurrency(currency string) *CoinbasePaymentMethodDeposit {
	opts.Currency = currency
	return opts
}
