package protomodel

import (
	"time"

	"github.com/alpine-hodler/sdk/serial"
)

// * This is a generated file, do not edit

// CoinbaseWebsocketTicker is real-time price updates every time a match
// happens. It batches updates in case of cascading matches, greatly reducing
// bandwidth requirements.
type CoinbaseWebsocketTicker struct {
	BestAsk   float64   `json:"best_ask"`
	BestBid   float64   `json:"best_bid"`
	LastSize  float64   `json:"last_size"`
	Price     float64   `json:"price"`
	ProductId string    `json:"product_id"`
	Sequence  int       `json:"sequence"`
	Side      string    `json:"side"`
	Time      time.Time `json:"time"`
	TradeId   int       `json:"trade_id"`
	Type      string    `json:"type"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseWebsocketTicker model
func (coinbaseWebsocketTicker *CoinbaseWebsocketTicker) UnmarshalJSON(d []byte) error {
	const (
		typeJsonTag      = "type"
		productIdJsonTag = "product_id"
		tradeIdJsonTag   = "trade_id"
		sequenceJsonTag  = "sequence"
		timeJsonTag      = "time"
		sideJsonTag      = "side"
		priceJsonTag     = "price"
		lastSizeJsonTag  = "last_size"
		bestBidJsonTag   = "best_bid"
		bestAskJsonTag   = "best_ask"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatFromString(bestAskJsonTag, &coinbaseWebsocketTicker.BestAsk)
	data.UnmarshalFloatFromString(bestBidJsonTag, &coinbaseWebsocketTicker.BestBid)
	data.UnmarshalFloatFromString(lastSizeJsonTag, &coinbaseWebsocketTicker.LastSize)
	data.UnmarshalFloatFromString(priceJsonTag, &coinbaseWebsocketTicker.Price)
	data.UnmarshalInt(sequenceJsonTag, &coinbaseWebsocketTicker.Sequence)
	data.UnmarshalInt(tradeIdJsonTag, &coinbaseWebsocketTicker.TradeId)
	data.UnmarshalString(productIdJsonTag, &coinbaseWebsocketTicker.ProductId)
	data.UnmarshalString(sideJsonTag, &coinbaseWebsocketTicker.Side)
	data.UnmarshalString(typeJsonTag, &coinbaseWebsocketTicker.Type)
	err = data.UnmarshalTime(time.RFC3339Nano, timeJsonTag, &coinbaseWebsocketTicker.Time)
	if err != nil {
		return err
	}
	return nil
}
