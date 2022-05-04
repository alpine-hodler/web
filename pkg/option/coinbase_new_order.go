package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/pkg/scalar"

// CoinbaseNewOrder are options for API requests.
type CoinbaseNewOrder struct {
	CancelAfter *scalar.CancelAfter `json:"cancel_after" bson:"cancel_after"`
	ClientOid   *string             `json:"client_oid" bson:"client_oid"`
	Funds       *float64            `json:"funds" bson:"funds"`
	PostOnly    *bool               `json:"post_only" bson:"post_only"`
	Price       *float64            `json:"price" bson:"price"`
	ProductId   string              `json:"product_id" bson:"product_id"`
	ProfileId   *string             `json:"profile_id" bson:"profile_id"`
	Side        scalar.OrderSide    `json:"side" bson:"side"`
	Size        *float64            `json:"size" bson:"size"`
	Stop        *scalar.OrderStop   `json:"stop" bson:"stop"`
	StopPrice   *float64            `json:"stop_price" bson:"stop_price"`
	Stp         *scalar.OrderSTP    `json:"stp" bson:"stp"`
	TimeInForce *scalar.TimeInForce `json:"time_in_force" bson:"time_in_force"`
	Type        *scalar.OrderType   `json:"type" bson:"type"`
}

// SetProfileId sets the ProfileId field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetProfileId(profileId string) *CoinbaseNewOrder {
	opts.ProfileId = &profileId
	return opts
}

// SetType sets the Type field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetType(typ scalar.OrderType) *CoinbaseNewOrder {
	opts.Type = &typ
	return opts
}

// SetSide sets the Side field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetSide(side scalar.OrderSide) *CoinbaseNewOrder {
	opts.Side = side
	return opts
}

// SetStp sets the Stp field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetStp(stp scalar.OrderSTP) *CoinbaseNewOrder {
	opts.Stp = &stp
	return opts
}

// SetStop sets the Stop field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetStop(stop scalar.OrderStop) *CoinbaseNewOrder {
	opts.Stop = &stop
	return opts
}

// SetStopPrice sets the StopPrice field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetStopPrice(stopPrice float64) *CoinbaseNewOrder {
	opts.StopPrice = &stopPrice
	return opts
}

// SetPrice sets the Price field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetPrice(price float64) *CoinbaseNewOrder {
	opts.Price = &price
	return opts
}

// SetSize sets the Size field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetSize(size float64) *CoinbaseNewOrder {
	opts.Size = &size
	return opts
}

// SetFunds sets the Funds field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetFunds(funds float64) *CoinbaseNewOrder {
	opts.Funds = &funds
	return opts
}

// SetProductId sets the ProductId field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetProductId(productId string) *CoinbaseNewOrder {
	opts.ProductId = productId
	return opts
}

// SetTimeInForce sets the TimeInForce field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetTimeInForce(timeInForce scalar.TimeInForce) *CoinbaseNewOrder {
	opts.TimeInForce = &timeInForce
	return opts
}

// SetCancelAfter sets the CancelAfter field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetCancelAfter(cancelAfter scalar.CancelAfter) *CoinbaseNewOrder {
	opts.CancelAfter = &cancelAfter
	return opts
}

// SetPostOnly sets the PostOnly field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetPostOnly(postOnly bool) *CoinbaseNewOrder {
	opts.PostOnly = &postOnly
	return opts
}

// SetClientOid sets the ClientOid field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetClientOid(clientOid string) *CoinbaseNewOrder {
	opts.ClientOid = &clientOid
	return opts
}
