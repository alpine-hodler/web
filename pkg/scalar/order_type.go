package scalar

type OrderType string

const (
	OrderTypeMarket OrderType = "market"
	OrderTypeLimit  OrderType = "limit"
)

// String converts a pointer reference to OrderType to a string
func (orderType *OrderType) String() (str string) {
	if orderType != nil {
		str = string(*orderType)
	}
	return
}
