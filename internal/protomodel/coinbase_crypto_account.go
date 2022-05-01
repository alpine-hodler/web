package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseCryptoAccount references a crypto account that a CoinbasePaymentMethod belongs to
type CoinbaseCryptoAccount struct {
	Id           string `json:"id" bson:"id"`
	Resource     string `json:"resource" bson:"resource"`
	ResourcePath string `json:"resource_path" bson:"resource_path"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseCryptoAccount model
func (coinbaseCryptoAccount *CoinbaseCryptoAccount) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag           = "id"
		resourceJsonTag     = "resource"
		resourcePathJsonTag = "resource_path"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(idJsonTag, &coinbaseCryptoAccount.Id)
	data.UnmarshalString(resourceJsonTag, &coinbaseCryptoAccount.Resource)
	data.UnmarshalString(resourcePathJsonTag, &coinbaseCryptoAccount.ResourcePath)
	return nil
}
