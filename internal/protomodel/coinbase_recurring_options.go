package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseRecurringOptions ??
type CoinbaseRecurringOptions struct {
	Label  string `json:"label" bson:"label"`
	Period string `json:"period" bson:"period"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseRecurringOptions model
func (coinbaseRecurringOptions *CoinbaseRecurringOptions) UnmarshalJSON(d []byte) error {
	const (
		periodJsonTag = "period"
		labelJsonTag  = "label"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(labelJsonTag, &coinbaseRecurringOptions.Label)
	data.UnmarshalString(periodJsonTag, &coinbaseRecurringOptions.Period)
	return nil
}
