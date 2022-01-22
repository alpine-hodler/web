package protomodel

import (
	"time"

	"github.com/alpine-hodler/sdk/internal/serial"
)

// * This is a generated file, do not edit

// CoinbaseProductTicker is a snapshot information about the last trade (tick),
// best bid/ask and 24h volume.
type CoinbaseProductTicker struct {
	Ask     float64   `json:"ask"`
	Bid     float64   `json:"bid"`
	Price   float64   `json:"price"`
	Size    float64   `json:"size"`
	Time    time.Time `json:"time"`
	TradeId int       `json:"trade_id"`
	Volume  float64   `json:"volume"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseProductTicker model
func (coinbaseProductTicker *CoinbaseProductTicker) UnmarshalJSON(d []byte) error {
	const (
		askJsonTag     = "ask"
		bidJsonTag     = "bid"
		volumeJsonTag  = "volume"
		tradeIdJsonTag = "trade_id"
		priceJsonTag   = "price"
		sizeJsonTag    = "size"
		timeJsonTag    = "time"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatFromString(askJsonTag, &coinbaseProductTicker.Ask)
	data.UnmarshalFloatFromString(bidJsonTag, &coinbaseProductTicker.Bid)
	data.UnmarshalFloatFromString(priceJsonTag, &coinbaseProductTicker.Price)
	data.UnmarshalFloatFromString(sizeJsonTag, &coinbaseProductTicker.Size)
	data.UnmarshalFloatFromString(volumeJsonTag, &coinbaseProductTicker.Volume)
	data.UnmarshalInt(tradeIdJsonTag, &coinbaseProductTicker.TradeId)
	err = data.UnmarshalTime(time.RFC3339Nano, timeJsonTag, &coinbaseProductTicker.Time)
	if err != nil {
		return err
	}
	return nil
}
