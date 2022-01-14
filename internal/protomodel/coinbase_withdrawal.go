package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseWithdrawal is data concerning withdrawing funds from the specified
// profile_id to a www.coinbase.com wallet.
type CoinbaseWithdrawal struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Fee      float64 `json:"fee"`
	Id       string  `json:"id"`
	PayoutAt string  `json:"payout_at"`
	Subtotal float64 `json:"subtotal"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseWithdrawal model
func (coinbaseWithdrawal *CoinbaseWithdrawal) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag       = "id"
		amountJsonTag   = "amount"
		currencyJsonTag = "currency"
		payoutAtJsonTag = "payout_at"
		feeJsonTag      = "fee"
		subtotalJsonTag = "subtotal"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatFromString(amountJsonTag, &coinbaseWithdrawal.Amount)
	data.UnmarshalFloatFromString(feeJsonTag, &coinbaseWithdrawal.Fee)
	data.UnmarshalFloatFromString(subtotalJsonTag, &coinbaseWithdrawal.Subtotal)
	data.UnmarshalString(currencyJsonTag, &coinbaseWithdrawal.Currency)
	data.UnmarshalString(idJsonTag, &coinbaseWithdrawal.Id)
	data.UnmarshalString(payoutAtJsonTag, &coinbaseWithdrawal.PayoutAt)
	return nil
}
