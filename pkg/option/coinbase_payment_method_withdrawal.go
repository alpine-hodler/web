package option

// * This is a generated file, do not edit

// CoinbasePaymentMethodWithdrawal are options for API requests.
type CoinbasePaymentMethodWithdrawal struct {
	Amount          float64 `json:"amount" bson:"amount"`
	Currency        string  `json:"currency" bson:"currency"`
	PaymentMethodId string  `json:"payment_method_id" bson:"payment_method_id"`
	ProfileId       *string `json:"profile_id" bson:"profile_id"`
}

// SetProfileId sets the ProfileId field on CoinbasePaymentMethodWithdrawal.
func (opts *CoinbasePaymentMethodWithdrawal) SetProfileId(profileId string) *CoinbasePaymentMethodWithdrawal {
	opts.ProfileId = &profileId
	return opts
}

// SetAmount sets the Amount field on CoinbasePaymentMethodWithdrawal.
func (opts *CoinbasePaymentMethodWithdrawal) SetAmount(amount float64) *CoinbasePaymentMethodWithdrawal {
	opts.Amount = amount
	return opts
}

// SetPaymentMethodId sets the PaymentMethodId field on CoinbasePaymentMethodWithdrawal.
func (opts *CoinbasePaymentMethodWithdrawal) SetPaymentMethodId(paymentMethodId string) *CoinbasePaymentMethodWithdrawal {
	opts.PaymentMethodId = paymentMethodId
	return opts
}

// SetCurrency sets the Currency field on CoinbasePaymentMethodWithdrawal.
func (opts *CoinbasePaymentMethodWithdrawal) SetCurrency(currency string) *CoinbasePaymentMethodWithdrawal {
	opts.Currency = currency
	return opts
}
