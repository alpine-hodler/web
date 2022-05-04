package option

// * This is a generated file, do not edit

// CoinbaseConversion are options for API requests.
type CoinbaseConversion struct {
	ProfileId *string `json:"profile_id" bson:"profile_id"`
}

// SetProfileId sets the ProfileId field on CoinbaseConversion.
func (opts *CoinbaseConversion) SetProfileId(profileId string) *CoinbaseConversion {
	opts.ProfileId = &profileId
	return opts
}
