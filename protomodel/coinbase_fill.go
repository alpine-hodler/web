package protomodel

import "github.com/alpine-hodler/sdk/serial"

// * This is a generated file, do not edit

// CoinbasePaymentMethod is a partial or complete match on a specific order.
type CoinbaseFill struct {
	Fee       float64 `json:"fee"`
	Liquidity string  `json:"liquidity"`
	OrderId   string  `json:"order_id"`
	Price     float64 `json:"price"`
	ProductId string  `json:"product_id"`
	ProfileId string  `json:"profile_id"`
	Settled   bool    `json:"settled"`
	Side      string  `json:"side"`
	Size      float64 `json:"size"`
	TradeId   int     `json:"trade_id"`
	UsdVolume float64 `json:"usd_volume"`
	UserId    string  `json:"user_id"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseFill model
func (coinbaseFill *CoinbaseFill) UnmarshalJSON(d []byte) error {
	const (
		tradeIdJsonTag   = "trade_id"
		productIdJsonTag = "product_id"
		orderIdJsonTag   = "order_id"
		userIdJsonTag    = "user_id"
		profileIdJsonTag = "profile_id"
		liquidityJsonTag = "liquidity"
		priceJsonTag     = "price"
		sizeJsonTag      = "size"
		feeJsonTag       = "fee"
		sideJsonTag      = "side"
		settledJsonTag   = "settled"
		usdVolumeJsonTag = "usd_volume"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(settledJsonTag, &coinbaseFill.Settled)
	data.UnmarshalFloatFromString(feeJsonTag, &coinbaseFill.Fee)
	data.UnmarshalFloatFromString(priceJsonTag, &coinbaseFill.Price)
	data.UnmarshalFloatFromString(sizeJsonTag, &coinbaseFill.Size)
	data.UnmarshalFloatFromString(usdVolumeJsonTag, &coinbaseFill.UsdVolume)
	data.UnmarshalInt(tradeIdJsonTag, &coinbaseFill.TradeId)
	data.UnmarshalString(liquidityJsonTag, &coinbaseFill.Liquidity)
	data.UnmarshalString(orderIdJsonTag, &coinbaseFill.OrderId)
	data.UnmarshalString(productIdJsonTag, &coinbaseFill.ProductId)
	data.UnmarshalString(profileIdJsonTag, &coinbaseFill.ProfileId)
	data.UnmarshalString(sideJsonTag, &coinbaseFill.Side)
	data.UnmarshalString(userIdJsonTag, &coinbaseFill.UserId)
	return nil
}
