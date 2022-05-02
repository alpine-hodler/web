package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseAccountDeposit are options for API requests.
type CoinbaseAccountDeposit struct {
	protomodel.CoinbaseAccountDepositOptions
}

// CoinbaseAccountDepositFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseAccountDepositFromPrototype(proto *protomodel.CoinbaseAccountDepositOptions) (opts *CoinbaseAccountDeposit) {
	if proto != nil {
		opts.CoinbaseAccountDepositOptions = *proto
	}
	return
}

// SetProfileID sets the ProfileID field on CoinbaseAccountDeposit.
func (opts *CoinbaseAccountDeposit) SetProfileID(profileId string) *CoinbaseAccountDeposit {
	opts.ProfileID = &profileId
	return opts
}

// SetAmount sets the Amount field on CoinbaseAccountDeposit.
func (opts *CoinbaseAccountDeposit) SetAmount(amount float64) *CoinbaseAccountDeposit {
	opts.Amount = amount
	return opts
}

// SetCoinbaseAccountID sets the CoinbaseAccountID field on CoinbaseAccountDeposit.
func (opts *CoinbaseAccountDeposit) SetCoinbaseAccountID(coinbaseAccountId string) *CoinbaseAccountDeposit {
	opts.CoinbaseAccountID = coinbaseAccountId
	return opts
}

// SetCurrency sets the Currency field on CoinbaseAccountDeposit.
func (opts *CoinbaseAccountDeposit) SetCurrency(currency string) *CoinbaseAccountDeposit {
	opts.Currency = currency
	return opts
}
