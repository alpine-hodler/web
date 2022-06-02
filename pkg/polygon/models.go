package polygon

// * This is a generated file, do not edit

// Upcoming returns information concerning market holidays and their open/close times.
type Upcoming struct {
	// Close is the market close time on the holiday (if it's not closed).
	Close string `json:"close" bson:"close"`

	// Date is the date of the holiday.
	Date string `json:"date" bson:"date"`

	// Exchange is market the record is for.
	Exchange string `json:"exchange" bson:"exchange"`

	// Name is the name of the holiday.
	Name string `json:"name" bson:"name"`

	// Open is the market open time on the holiday (if it's not closed).
	Open string `json:"open" bson:"open"`

	// Status is the status of the market on the holiday.
	Status string `json:"status" bson:"status"`
}
