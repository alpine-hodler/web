package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseAccountTransfers are options for API requests.
type CoinbaseAccountTransfers struct {
	protomodel.CoinbaseAccountTransfersOptions
}

// CoinbaseAccountTransfersFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseAccountTransfersFromPrototype(proto *protomodel.CoinbaseAccountTransfersOptions) (opts *CoinbaseAccountTransfers) {
	if proto != nil {
		opts.CoinbaseAccountTransfersOptions = *proto
	}
	return
}

// SetBefore sets the Before field on CoinbaseAccountTransfers.
func (opts *CoinbaseAccountTransfers) SetBefore(before string) *CoinbaseAccountTransfers {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on CoinbaseAccountTransfers.
func (opts *CoinbaseAccountTransfers) SetAfter(after string) *CoinbaseAccountTransfers {
	opts.After = &after
	return opts
}

// SetLimit sets the Limit field on CoinbaseAccountTransfers.
func (opts *CoinbaseAccountTransfers) SetLimit(limit int) *CoinbaseAccountTransfers {
	opts.Limit = &limit
	return opts
}

// SetType sets the Type field on CoinbaseAccountTransfers.
func (opts *CoinbaseAccountTransfers) SetType(t string) *CoinbaseAccountTransfers {
	opts.Type = &t
	return opts
}
