package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseAccountTransferDetails are the details for an account transfer.
type CoinbaseAccountTransferDetails struct {
	CoinbaseAccountId       string `json:"coinbase_account_id" bson:"coinbase_account_id"`
	CoinbasePaymentMethodId string `json:"coinbase_payment_method_id" bson:"coinbase_payment_method_id"`
	CoinbaseTransactionId   string `json:"coinbase_transaction_id" bson:"coinbase_transaction_id"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseAccountTransferDetails model
func (coinbaseAccountTransferDetails *CoinbaseAccountTransferDetails) UnmarshalJSON(d []byte) error {
	const (
		coinbaseAccountIdJsonTag       = "coinbase_account_id"
		coinbaseTransactionIdJsonTag   = "coinbase_transaction_id"
		coinbasePaymentMethodIdJsonTag = "coinbase_payment_method_id"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(coinbaseAccountIdJsonTag, &coinbaseAccountTransferDetails.CoinbaseAccountId)
	data.UnmarshalString(coinbasePaymentMethodIdJsonTag, &coinbaseAccountTransferDetails.CoinbasePaymentMethodId)
	data.UnmarshalString(coinbaseTransactionIdJsonTag, &coinbaseAccountTransferDetails.CoinbaseTransactionId)
	return nil
}
