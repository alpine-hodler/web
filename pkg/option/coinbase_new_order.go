package option

// * This is a generated file, do not edit
import (
	"github.com/alpine-hodler/sdk/internal/protomodel"
	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// CoinbaseNewOrder are options for API requests.
type CoinbaseNewOrder struct {
	protomodel.CoinbaseNewOrderOptions
}

// CoinbaseNewOrderFromPrototype will initialize the exportable options struct from the internal prototype.
func CoinbaseNewOrderFromPrototype(proto *protomodel.CoinbaseNewOrderOptions) (opts *CoinbaseNewOrder) {
	if proto != nil {
		opts.CoinbaseNewOrderOptions = *proto
	}
	return
}

// SetProfileID sets the ProfileID field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetProfileID(profileId string) *CoinbaseNewOrder {
	opts.ProfileID = &profileId
	return opts
}

// SetType sets the Type field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetType(t scalar.OrderType) *CoinbaseNewOrder {
	opts.Type = &t
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

// SetProductID sets the ProductID field on CoinbaseNewOrder.
func (opts *CoinbaseNewOrder) SetProductID(productId string) *CoinbaseNewOrder {
	opts.ProductID = productId
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
