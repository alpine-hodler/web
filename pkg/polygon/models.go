package polygon

// * This is a generated file, do not edit

// Upcoming returns information concerning market holidays and their open/close times.
type Upcoming struct {
	Close    string `json:"close" bson:"close"`
	Date     string `json:"date" bson:"date"`
	Exchange string `json:"exchange" bson:"exchange"`
	Name     string `json:"name" bson:"name"`
	Open     string `json:"open" bson:"open"`
	Status   string `json:"status" bson:"status"`
}
