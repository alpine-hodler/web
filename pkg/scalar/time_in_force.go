package scalar

// OrderTimeInForce (Time in force) policies provide guarantees about the
// lifetime of an order
type TimeInForce string

const (
	// GTC Good till canceled orders remain open on the book until canceled. This
	// is the default behavior if no policy is specified.
	TimeInForceGTC = "GTC"

	// IOC Immediate or cancel orders instantly cancel the remaining size of the
	// limit order instead of opening it on the book.
	TimeInForceIOC = "IOC"

	// FOK Fill or kill orders are rejected if the entire size cannot be matched.
	TimeInForceFOK = "FOK"

	// GTT Good till time orders remain open on the book until canceled or the
	// allotted cancel_after is depleted on the matching engine. GTT orders are
	// guaranteed to cancel before any other order is processed after the
	// cancel_after timestamp which is returned by the API. A day is considered 24
	// hours.
	TimeInForceGTT = "GTT"
)

func (timeInForce *TimeInForce) String() (str string) {
	if timeInForce != nil {
		str = string(*timeInForce)
	}
	return
}
