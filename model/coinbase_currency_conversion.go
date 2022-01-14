package model

// * This file was initialized by the meta-program, but is open to modification
import "github.com/alpine-hodler/sdk/protomodel"

// CoinbaseCurrencyConversion is the response that converts funds from from
// currency to to currency. Funds are converted on the from account in the
// profile_id profile.
type CoinbaseCurrencyConversion struct {
	protomodel.CoinbaseCurrencyConversion
}
