package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// CoinbaseCryptoAddressInfo holds info for a crypto address
type CoinbaseCryptoAddressInfo struct {
	Address        string `json:"address" bson:"address"`
	DestinationTag string `json:"destination_tag" bson:"destination_tag"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseCryptoAddressInfo model
func (coinbaseCryptoAddressInfo *CoinbaseCryptoAddressInfo) UnmarshalJSON(d []byte) error {
	const (
		addressJsonTag        = "address"
		destinationTagJsonTag = "destination_tag"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(addressJsonTag, &coinbaseCryptoAddressInfo.Address)
	data.UnmarshalString(destinationTagJsonTag, &coinbaseCryptoAddressInfo.DestinationTag)
	return nil
}
