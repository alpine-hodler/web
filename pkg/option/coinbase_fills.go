package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseFills are options for API requests.
type CoinbaseFills struct {
	protomodel.CoinbaseFillsOptions
}

// CoinbaseFillsFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseFillsFromPrototype(proto *protomodel.CoinbaseFillsOptions) (opts *CoinbaseFills) {
	if proto != nil {
		opts.CoinbaseFillsOptions = *proto
	}
	return
}

// SetOrderID sets the OrderID field on CoinbaseFills.
func (opts *CoinbaseFills) SetOrderID(orderId string) *CoinbaseFills {
	opts.OrderID = &orderId
	return opts
}

// SetProductID sets the ProductID field on CoinbaseFills.
func (opts *CoinbaseFills) SetProductID(productId string) *CoinbaseFills {
	opts.ProductID = &productId
	return opts
}

// SetProfileID sets the ProfileID field on CoinbaseFills.
func (opts *CoinbaseFills) SetProfileID(profileId string) *CoinbaseFills {
	opts.ProfileID = &profileId
	return opts
}

// SetLimit sets the Limit field on CoinbaseFills.
func (opts *CoinbaseFills) SetLimit(limit int) *CoinbaseFills {
	opts.Limit = &limit
	return opts
}

// SetBefore sets the Before field on CoinbaseFills.
func (opts *CoinbaseFills) SetBefore(before int) *CoinbaseFills {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on CoinbaseFills.
func (opts *CoinbaseFills) SetAfter(after int) *CoinbaseFills {
	opts.After = &after
	return opts
}
