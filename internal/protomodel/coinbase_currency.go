package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseCurrency is a currency that coinbase knows about. Not al currencies may be currently in use for trading.
type CoinbaseCurrency struct {
	ConvertibleTo []string                `json:"convertible_to" bson:"convertible_to"`
	Id            string                  `json:"id" bson:"id"`
	MaxPrecision  float64                 `json:"max_precision" bson:"max_precision"`
	Message       string                  `json:"message" bson:"message"`
	MinSize       float64                 `json:"min_size" bson:"min_size"`
	Name          string                  `json:"name" bson:"name"`
	ProtoDetails  CoinbaseCurrencyDetails `json:"details" bson:"details"`
	Status        string                  `json:"status" bson:"status"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseCurrency model
func (coinbaseCurrency *CoinbaseCurrency) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag            = "id"
		nameJsonTag          = "name"
		minSizeJsonTag       = "min_size"
		statusJsonTag        = "status"
		messageJsonTag       = "message"
		maxPrecisionJsonTag  = "max_precision"
		convertibleToJsonTag = "convertible_to"
		detailsJsonTag       = "details"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbaseCurrency.ProtoDetails = CoinbaseCurrencyDetails{}
	if err := data.UnmarshalStruct(detailsJsonTag, &coinbaseCurrency.ProtoDetails); err != nil {
		return err
	}
	data.UnmarshalFloatFromString(maxPrecisionJsonTag, &coinbaseCurrency.MaxPrecision)
	data.UnmarshalFloatFromString(minSizeJsonTag, &coinbaseCurrency.MinSize)
	data.UnmarshalString(idJsonTag, &coinbaseCurrency.Id)
	data.UnmarshalString(messageJsonTag, &coinbaseCurrency.Message)
	data.UnmarshalString(nameJsonTag, &coinbaseCurrency.Name)
	data.UnmarshalString(statusJsonTag, &coinbaseCurrency.Status)
	data.UnmarshalStringSlice(convertibleToJsonTag, &coinbaseCurrency.ConvertibleTo)
	return nil
}
