package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseAccountWithdrawal are options for API requests.
type CoinbaseAccountWithdrawal struct {
	protomodel.CoinbaseAccountWithdrawalOptions
}

// CoinbaseAccountWithdrawalFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseAccountWithdrawalFromPrototype(proto *protomodel.CoinbaseAccountWithdrawalOptions) (opts *CoinbaseAccountWithdrawal) {
	if proto != nil {
		opts.CoinbaseAccountWithdrawalOptions = *proto
	}
	return
}

// SetProfileID sets the ProfileID field on CoinbaseAccountWithdrawal.
func (opts *CoinbaseAccountWithdrawal) SetProfileID(profileId string) *CoinbaseAccountWithdrawal {
	opts.ProfileID = &profileId
	return opts
}

// SetAmount sets the Amount field on CoinbaseAccountWithdrawal.
func (opts *CoinbaseAccountWithdrawal) SetAmount(amount float64) *CoinbaseAccountWithdrawal {
	opts.Amount = amount
	return opts
}

// SetCoinbaseAccountID sets the CoinbaseAccountID field on CoinbaseAccountWithdrawal.
func (opts *CoinbaseAccountWithdrawal) SetCoinbaseAccountID(coinbaseAccountId string) *CoinbaseAccountWithdrawal {
	opts.CoinbaseAccountID = coinbaseAccountId
	return opts
}

// SetCurrency sets the Currency field on CoinbaseAccountWithdrawal.
func (opts *CoinbaseAccountWithdrawal) SetCurrency(currency string) *CoinbaseAccountWithdrawal {
	opts.Currency = currency
	return opts
}
