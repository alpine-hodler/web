package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbasePaymentMethodDeposit are options for API requests.
type CoinbasePaymentMethodDeposit struct {
	protomodel.CoinbasePaymentMethodDepositOptions
}

// CoinbasePaymentMethodDepositFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbasePaymentMethodDepositFromPrototype(proto *protomodel.CoinbasePaymentMethodDepositOptions) (opts *CoinbasePaymentMethodDeposit) {
	if proto != nil {
		opts.CoinbasePaymentMethodDepositOptions = *proto
	}
	return
}

// SetProfileID sets the ProfileID field on CoinbasePaymentMethodDeposit.
func (opts *CoinbasePaymentMethodDeposit) SetProfileID(profileId string) *CoinbasePaymentMethodDeposit {
	opts.ProfileID = &profileId
	return opts
}

// SetAmount sets the Amount field on CoinbasePaymentMethodDeposit.
func (opts *CoinbasePaymentMethodDeposit) SetAmount(amount float64) *CoinbasePaymentMethodDeposit {
	opts.Amount = amount
	return opts
}

// SetPaymentMethodID sets the PaymentMethodID field on CoinbasePaymentMethodDeposit.
func (opts *CoinbasePaymentMethodDeposit) SetPaymentMethodID(paymentMethodId string) *CoinbasePaymentMethodDeposit {
	opts.PaymentMethodID = paymentMethodId
	return opts
}

// SetCurrency sets the Currency field on CoinbasePaymentMethodDeposit.
func (opts *CoinbasePaymentMethodDeposit) SetCurrency(currency string) *CoinbasePaymentMethodDeposit {
	opts.Currency = currency
	return opts
}
