package scalar

// OrderCancelTime represents when to cancel an order
type OrderCancelTime string

const (
	OrderCancelTimeMin  = "min"
	OrderCancelTimeHour = "hour"
	OrderCancelTimeDay  = "day"
)
