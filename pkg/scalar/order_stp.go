package scalar

// OrderSTP is the order's Self trade preservation flag
type OrderSTP string

const (
	// OrderSTP_DC = decrease and cancel, default
	OrderSTP_DC = "dc"

	// OrderSTP_CO = cancel oldest
	OrderSTP_CO OrderSTP = "co"

	// OrderSTP_CN = cancel newest
	OrderSTP_CN OrderSTP = "cn"

	// OrderSTP_CB = cancel boath
	OrderSTP_CB OrderSTP = "cb"
)

func (orderSTP *OrderSTP) String() (str string) {
	if orderSTP != nil {
		str = string(*orderSTP)
	}
	return
}
