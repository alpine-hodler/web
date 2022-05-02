package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbasePaymentMethodWithdrawal are options for API requests.
type CoinbasePaymentMethodWithdrawal struct {
	protomodel.CoinbasePaymentMethodWithdrawalOptions
}

// CoinbasePaymentMethodWithdrawalFromPrototype will initialize the exportable options struct from the internal
// prototype.
func CoinbasePaymentMethodWithdrawalFromPrototype(proto *protomodel.CoinbasePaymentMethodWithdrawalOptions) (opts *CoinbasePaymentMethodWithdrawal) {
	if proto != nil {
		opts.CoinbasePaymentMethodWithdrawalOptions = *proto
	}
	return
}

// SetProfileID sets the ProfileID field on CoinbasePaymentMethodWithdrawal.
func (opts *CoinbasePaymentMethodWithdrawal) SetProfileID(profileId string) *CoinbasePaymentMethodWithdrawal {
	opts.ProfileID = &profileId
	return opts
}

// SetAmount sets the Amount field on CoinbasePaymentMethodWithdrawal.
func (opts *CoinbasePaymentMethodWithdrawal) SetAmount(amount float64) *CoinbasePaymentMethodWithdrawal {
	opts.Amount = amount
	return opts
}

// SetPaymentMethodID sets the PaymentMethodID field on CoinbasePaymentMethodWithdrawal.
func (opts *CoinbasePaymentMethodWithdrawal) SetPaymentMethodID(paymentMethodId string) *CoinbasePaymentMethodWithdrawal {
	opts.PaymentMethodID = paymentMethodId
	return opts
}

// SetCurrency sets the Currency field on CoinbasePaymentMethodWithdrawal.
func (opts *CoinbasePaymentMethodWithdrawal) SetCurrency(currency string) *CoinbasePaymentMethodWithdrawal {
	opts.Currency = currency
	return opts
}
