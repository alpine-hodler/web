package option

// * This is a generated file, do not edit

// CoinbaseAccountLedger are options for API requests.
type CoinbaseAccountLedger struct {
	After     *string `json:"after" bson:"after"`
	Before    *string `json:"before" bson:"before"`
	EndDate   *string `json:"end_date" bson:"end_date"`
	Limit     *int    `json:"limit" bson:"limit"`
	ProfileId *string `json:"profile_id" bson:"profile_id"`
	StartDate *string `json:"start_date" bson:"start_date"`
}

// SetStartDate sets the StartDate field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetStartDate(startDate string) *CoinbaseAccountLedger {
	opts.StartDate = &startDate
	return opts
}

// SetEndDate sets the EndDate field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetEndDate(endDate string) *CoinbaseAccountLedger {
	opts.EndDate = &endDate
	return opts
}

// SetBefore sets the Before field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetBefore(before string) *CoinbaseAccountLedger {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetAfter(after string) *CoinbaseAccountLedger {
	opts.After = &after
	return opts
}

// SetProfileId sets the ProfileId field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetProfileId(profileId string) *CoinbaseAccountLedger {
	opts.ProfileId = &profileId
	return opts
}

// SetLimit sets the Limit field on CoinbaseAccountLedger.
func (opts *CoinbaseAccountLedger) SetLimit(limit int) *CoinbaseAccountLedger {
	opts.Limit = &limit
	return opts
}
