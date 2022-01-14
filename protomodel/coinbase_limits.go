package protomodel

import "github.com/alpine-hodler/sdk/serial"

// * This is a generated file, do not edit

// CoinbaseLimits defines limits for a payment method
type CoinbaseLimits struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseLimits model
func (coinbaseLimits *CoinbaseLimits) UnmarshalJSON(d []byte) error {
	const (
		typeJsonTag = "type"
		nameJsonTag = "name"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(nameJsonTag, &coinbaseLimits.Name)
	data.UnmarshalString(typeJsonTag, &coinbaseLimits.Type)
	return nil
}
