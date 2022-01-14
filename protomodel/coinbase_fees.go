package protomodel

import "github.com/alpine-hodler/sdk/serial"

// * This is a generated file, do not edit

// CoinbaseFees are fees rates and 30 days trailing volume.
type CoinbaseFees struct {
	MakerFeeRate float64 `json:"maker_fee_rate"`
	TakerFeeRate float64 `json:"taker_fee_rate"`
	UsdVolume    float64 `json:"usd_volume"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseFees model
func (coinbaseFees *CoinbaseFees) UnmarshalJSON(d []byte) error {
	const (
		takerFeeRateJsonTag = "taker_fee_rate"
		makerFeeRateJsonTag = "maker_fee_rate"
		usdVolumeJsonTag    = "usd_volume"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatFromString(makerFeeRateJsonTag, &coinbaseFees.MakerFeeRate)
	data.UnmarshalFloatFromString(takerFeeRateJsonTag, &coinbaseFees.TakerFeeRate)
	data.UnmarshalFloatFromString(usdVolumeJsonTag, &coinbaseFees.UsdVolume)
	return nil
}
