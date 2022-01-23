package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseWithdrawalFeeEstimate is a fee estimate for the crypto withdrawal to
// crypto address
type CoinbaseWithdrawalFeeEstimate struct {
	Fee float64 `json:"fee" bson:"fee"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseWithdrawalFeeEstimate
// model
func (coinbaseWithdrawalFeeEstimate *CoinbaseWithdrawalFeeEstimate) UnmarshalJSON(d []byte) error {
	const (
		feeJsonTag = "fee"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloat(feeJsonTag, &coinbaseWithdrawalFeeEstimate.Fee)
	return nil
}
