package scalar

// OrderStop is either loss or entry
type OrderStop string

const (
	OrderStopLoss  = "loss"
	OrderStopEntry = "entry"
)

func (orderStop *OrderStop) String() (str string) {
	if orderStop != nil {
		str = string(*orderStop)
	}
	return
}
