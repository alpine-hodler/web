package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseCryptoAddressWarning is a warning for generating a crypting address
type CoinbaseCryptoAddressWarning struct {
	Details  string `json:"details"`
	ImageUrl string `json:"image_url"`
	Title    string `json:"title"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseCryptoAddressWarning
// model
func (coinbaseCryptoAddressWarning *CoinbaseCryptoAddressWarning) UnmarshalJSON(d []byte) error {
	const (
		titleJsonTag    = "title"
		detailsJsonTag  = "details"
		imageUrlJsonTag = "image_url"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(detailsJsonTag, &coinbaseCryptoAddressWarning.Details)
	data.UnmarshalString(imageUrlJsonTag, &coinbaseCryptoAddressWarning.ImageUrl)
	data.UnmarshalString(titleJsonTag, &coinbaseCryptoAddressWarning.Title)
	return nil
}
