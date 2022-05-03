package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseProducts are options for API requests.
type CoinbaseProducts struct {
	protomodel.CoinbaseProductsOptions
}

// CoinbaseProductsFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseProductsFromPrototype(proto *protomodel.CoinbaseProductsOptions) (opts *CoinbaseProducts) {
	if proto != nil {
		opts.CoinbaseProductsOptions = *proto
	}
	return
}

// SetType sets the Type field on CoinbaseProducts.
func (opts *CoinbaseProducts) SetType(t string) *CoinbaseProducts {
	opts.Type = &t
	return opts
}
