package scalar

// OrderSide represents "Buy" or "Sell"
type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

func (orderSide *OrderSide) String() (str string) {
	if orderSide != nil {
		str = string(*orderSide)
	}
	return
}
