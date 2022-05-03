package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseWithdrawalFeeEstimate are options for API requests.
type CoinbaseWithdrawalFeeEstimate struct {
	protomodel.CoinbaseWithdrawalFeeEstimateOptions
}

// CoinbaseWithdrawalFeeEstimateFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseWithdrawalFeeEstimateFromPrototype(proto *protomodel.CoinbaseWithdrawalFeeEstimateOptions) (opts *CoinbaseWithdrawalFeeEstimate) {
	if proto != nil {
		opts.CoinbaseWithdrawalFeeEstimateOptions = *proto
	}
	return
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
