package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseAccountHolds are options for API requests.
type CoinbaseAccountHolds struct {
	protomodel.CoinbaseAccountHoldsOptions
}

// CoinbaseAccountHoldsFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseAccountHoldsFromPrototype(proto *protomodel.CoinbaseAccountHoldsOptions) (opts *CoinbaseAccountHolds) {
	if proto != nil {
		opts.CoinbaseAccountHoldsOptions = *proto
	}
	return
}

// SetBefore sets the Before field on CoinbaseAccountHolds.
func (opts *CoinbaseAccountHolds) SetBefore(before string) *CoinbaseAccountHolds {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on CoinbaseAccountHolds.
func (opts *CoinbaseAccountHolds) SetAfter(after string) *CoinbaseAccountHolds {
	opts.After = &after
	return opts
}

// SetLimit sets the Limit field on CoinbaseAccountHolds.
func (opts *CoinbaseAccountHolds) SetLimit(limit int) *CoinbaseAccountHolds {
	opts.Limit = &limit
	return opts
}
