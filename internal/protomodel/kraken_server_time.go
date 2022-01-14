package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// KrakenServerTime holds data concerning the server time
type KrakenServerTime struct {
	Error       []string               `json:"error"`
	ProtoResult KrakenServerTimeResult `json:"result"`
}

// UnmarshalJSON will deserialize bytes into a KrakenServerTime model
func (krakenServerTime *KrakenServerTime) UnmarshalJSON(d []byte) error {
	const (
		resultJsonTag = "result"
		errorJsonTag  = "error"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalStringSlice(errorJsonTag, &krakenServerTime.Error)
	krakenServerTime.ProtoResult = KrakenServerTimeResult{}
	if err := data.UnmarshalStruct(resultJsonTag, &krakenServerTime.ProtoResult); err != nil {
		return err
	}
	return nil
}
