package scalar

// OrderStop is either loss or entry
type OrderStop string

const (
	OrderStopLoss  OrderStop = "loss"
	OrderStopEntry OrderStop = "entry"
)

func (orderStop *OrderStop) String() (str string) {
	if orderStop != nil {
		str = string(*orderStop)
	}
	return
}
