package scalar

// OrderCancelTime represents when to cancel an order
type OrderCancelTime string

const (
	OrderCancelTimeMin  OrderCancelTime = "min"
	OrderCancelTimeHour OrderCancelTime = "hour"
	OrderCancelTimeDay  OrderCancelTime = "day"
)
