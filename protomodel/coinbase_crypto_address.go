package protomodel

import (
	"encoding/json"
	"time"

	"github.com/alpine-hodler/sdk/serial"
)

// * This is a generated file, do not edit

// CoinbaseCryptoAddress is used for a one-time crypto address for depositing
// crypto.
type CoinbaseCryptoAddress struct {
	Address          string                          `json:"address"`
	CallbackUrl      string                          `json:"callback_url"`
	CreateAt         time.Time                       `json:"create_at"`
	DepositUri       string                          `json:"deposit_uri"`
	DestinationTag   string                          `json:"destination_tag"`
	Id               string                          `json:"id"`
	LegacyAddress    string                          `json:"legacy_address"`
	Name             string                          `json:"name"`
	Network          string                          `json:"network"`
	ProtoAddressInfo CoinbaseCryptoAddressInfo       `json:"address_info"`
	ProtoWarnings    []*CoinbaseCryptoAddressWarning `json:"warnings"`
	Resource         string                          `json:"resource"`
	ResourcePath     string                          `json:"resource_path"`
	UpdatedAt        time.Time                       `json:"updated_at"`
	UriScheme        string                          `json:"uri_scheme"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseCryptoAddress model
func (coinbaseCryptoAddress *CoinbaseCryptoAddress) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag             = "id"
		addressJsonTag        = "address"
		addressInfoJsonTag    = "address_info"
		nameJsonTag           = "name"
		createAtJsonTag       = "create_at"
		updatedAtJsonTag      = "updated_at"
		networkJsonTag        = "network"
		uriSchemeJsonTag      = "uri_scheme"
		resourceJsonTag       = "resource"
		resourcePathJsonTag   = "resource_path"
		warningsJsonTag       = "warnings"
		legacyAddressJsonTag  = "legacy_address"
		destinationTagJsonTag = "destination_tag"
		depositUriJsonTag     = "deposit_uri"
		callbackUrlJsonTag    = "callback_url"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbaseCryptoAddress.ProtoAddressInfo = CoinbaseCryptoAddressInfo{}
	if err := data.UnmarshalStruct(addressInfoJsonTag, &coinbaseCryptoAddress.ProtoAddressInfo); err != nil {
		return err
	}
	data.UnmarshalString(addressJsonTag, &coinbaseCryptoAddress.Address)
	data.UnmarshalString(callbackUrlJsonTag, &coinbaseCryptoAddress.CallbackUrl)
	data.UnmarshalString(depositUriJsonTag, &coinbaseCryptoAddress.DepositUri)
	data.UnmarshalString(destinationTagJsonTag, &coinbaseCryptoAddress.DestinationTag)
	data.UnmarshalString(idJsonTag, &coinbaseCryptoAddress.Id)
	data.UnmarshalString(legacyAddressJsonTag, &coinbaseCryptoAddress.LegacyAddress)
	data.UnmarshalString(nameJsonTag, &coinbaseCryptoAddress.Name)
	data.UnmarshalString(networkJsonTag, &coinbaseCryptoAddress.Network)
	data.UnmarshalString(resourceJsonTag, &coinbaseCryptoAddress.Resource)
	data.UnmarshalString(resourcePathJsonTag, &coinbaseCryptoAddress.ResourcePath)
	data.UnmarshalString(uriSchemeJsonTag, &coinbaseCryptoAddress.UriScheme)
	err = data.UnmarshalTime(time.RFC3339Nano, createAtJsonTag, &coinbaseCryptoAddress.CreateAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, updatedAtJsonTag, &coinbaseCryptoAddress.UpdatedAt)
	if err != nil {
		return err
	}
	if v := data.Value(warningsJsonTag); v != nil {
		for _, item := range data.Value(warningsJsonTag).([]interface{}) {
			bytes, _ := json.Marshal(item)
			obj := CoinbaseCryptoAddressWarning{}
			if err := json.Unmarshal(bytes, &obj); err != nil {
				return err
			}
			coinbaseCryptoAddress.ProtoWarnings = append(coinbaseCryptoAddress.ProtoWarnings, &obj)
		}
	}
	return nil
}
