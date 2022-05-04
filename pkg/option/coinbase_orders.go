package option

// * This is a generated file, do not edit
import "time"

// CoinbaseOrders are options for API requests.
type CoinbaseOrders struct {
	After     *string    `json:"after" bson:"after"`
	Before    *string    `json:"before" bson:"before"`
	EndDate   *time.Time `json:"end_date" bson:"end_date"`
	Limit     int        `json:"limit" bson:"limit"`
	ProductId *string    `json:"product_id" bson:"product_id"`
	ProfileId *string    `json:"profile_id" bson:"profile_id"`
	SortedBy  *string    `json:"sortedBy" bson:"sortedBy"`
	Sorting   *string    `json:"sorting" bson:"sorting"`
	StartDate *time.Time `json:"start_date" bson:"start_date"`
	Status    []string   `json:"status" bson:"status"`
}

// SetProfileId sets the ProfileId field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetProfileId(profileId string) *CoinbaseOrders {
	opts.ProfileId = &profileId
	return opts
}

// SetProductId sets the ProductId field on CoinbaseOrders.
func (opts *CoinbaseOrders) SetProductId(productId string) *CoinbaseOrders {
	opts.ProductId = &productId
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
func (opts *CoinbaseOrders) SetStatus(status []string) *CoinbaseOrders {
	opts.Status = status
	return opts
}
