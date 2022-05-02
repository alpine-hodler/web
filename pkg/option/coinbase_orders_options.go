package option

// * This is a generated file, do not edit
import (
	"time"

	"github.com/alpine-hodler/sdk/internal/protomodel"
)

// CoinbaseOrders are options for API requests.
type CoinbaseOrders struct {
	protomodel.CoinbaseOrdersOptions
}

// CoinbaseOrdersFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseOrdersFromPrototype(proto *protomodel.CoinbaseOrdersOptions) (opts *CoinbaseOrders) {
	if proto != nil {
		opts.CoinbaseOrdersOptions = *proto
	}
	return
}

// SetProfileID sets the ProfileID field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetProfileID(profileId string) *CoinbaseOrders {
	opts.ProfileID = &profileId
	return opts
}

// SetProductID sets the ProductID field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetProductID(productId string) *CoinbaseOrders {
	opts.ProductID = &productId
	return opts
}

// SetSortedBy sets the SortedBy field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetSortedBy(sortedBy string) *CoinbaseOrders {
	opts.SortedBy = &sortedBy
	return opts
}

// SetSorting sets the Sorting field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetSorting(sorting string) *CoinbaseOrders {
	opts.Sorting = &sorting
	return opts
}

// SetStartDate sets the StartDate field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetStartDate(startDate time.Time) *CoinbaseOrders {
	opts.StartDate = &startDate
	return opts
}

// SetEndDate sets the EndDate field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetEndDate(endDate time.Time) *CoinbaseOrders {
	opts.EndDate = &endDate
	return opts
}

// SetBefore sets the Before field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetBefore(before string) *CoinbaseOrders {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetAfter(after string) *CoinbaseOrders {
	opts.After = &after
	return opts
}

// SetLimit sets the Limit field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetLimit(limit int) *CoinbaseOrders {
	opts.Limit = limit
	return opts
}

// SetStatus sets the Status field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetStatus(status []*string) *CoinbaseOrders {
	opts.Status = status
	return opts
}
