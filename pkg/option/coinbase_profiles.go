package option

// * This is a generated file, do not edit

// CoinbaseProfiles are options for API requests.
type CoinbaseProfiles struct {
	Active *bool `json:"active" bson:"active"`
}

// SetActive sets the Active field on CoinbaseProfiles.
func (opts *CoinbaseProfiles) SetActive(active bool) *CoinbaseProfiles {
	opts.Active = &active
	return opts
}
