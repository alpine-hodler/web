package scalar

// CancelAfter is the timeframe in which to cancel an order if it hasn't been
// filled
type CancelAfter string

const (
	CancelAfterMin  = "min"
	CancelAfterHour = "hour"
	CancelAfterDay  = "day"
)

func (cancelAfter *CancelAfter) String() (str string) {
	if cancelAfter != nil {
		str = string(*cancelAfter)
	}
	return
}
