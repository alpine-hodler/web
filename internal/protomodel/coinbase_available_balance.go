package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseAvailableBalance is the available balance on the coinbase account
type CoinbaseAvailableBalance struct {
	Amount   float64 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
	Scale    string  `json:"scale" bson:"scale"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseAvailableBalance model
func (coinbaseAvailableBalance *CoinbaseAvailableBalance) UnmarshalJSON(d []byte) error {
	const (
		amountJsonTag   = "amount"
		currencyJsonTag = "currency"
		scaleJsonTag    = "scale"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatFromString(amountJsonTag, &coinbaseAvailableBalance.Amount)
	data.UnmarshalString(currencyJsonTag, &coinbaseAvailableBalance.Currency)
	data.UnmarshalString(scaleJsonTag, &coinbaseAvailableBalance.Scale)
	return nil
}
