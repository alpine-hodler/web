package protomodel

import "github.com/alpine-hodler/sdk/serial"

// * This is a generated file, do not edit

// CoinbaseAccountLedgerDetails are the details for account history.
type CoinbaseAccountLedgerDetails struct {
	OrderId   string `json:"order_id"`
	ProductId string `json:"product_id"`
	TradeId   string `json:"trade_id"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseAccountLedgerDetails
// model
func (coinbaseAccountLedgerDetails *CoinbaseAccountLedgerDetails) UnmarshalJSON(d []byte) error {
	const (
		orderIdJsonTag   = "order_id"
		tradeIdJsonTag   = "trade_id"
		productIdJsonTag = "product_id"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(orderIdJsonTag, &coinbaseAccountLedgerDetails.OrderId)
	data.UnmarshalString(productIdJsonTag, &coinbaseAccountLedgerDetails.ProductId)
	data.UnmarshalString(tradeIdJsonTag, &coinbaseAccountLedgerDetails.TradeId)
	return nil
}
