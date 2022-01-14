package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseBalance is the balance for picker data
type CoinbaseBalance struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseBalance model
func (coinbaseBalance *CoinbaseBalance) UnmarshalJSON(d []byte) error {
	const (
		amountJsonTag   = "amount"
		currencyJsonTag = "currency"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatFromString(amountJsonTag, &coinbaseBalance.Amount)
	data.UnmarshalString(currencyJsonTag, &coinbaseBalance.Currency)
	return nil
}
