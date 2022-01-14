package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseBankCountry are the name and code for the bank's country associated
// with a wallet
type CoinbaseBankCountry struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseBankCountry model
func (coinbaseBankCountry *CoinbaseBankCountry) UnmarshalJSON(d []byte) error {
	const (
		nameJsonTag = "name"
		codeJsonTag = "code"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(codeJsonTag, &coinbaseBankCountry.Code)
	data.UnmarshalString(nameJsonTag, &coinbaseBankCountry.Name)
	return nil
}
