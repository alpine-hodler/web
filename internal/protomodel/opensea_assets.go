package protomodel

import (
	"encoding/json"

	"github.com/alpine-hodler/sdk/internal/serial"
)

// * This is a generated file, do not edit

// OpenseaAssets are a set of assets from opensea NFTs
type OpenseaAssets struct {
	ProtoAssets []*OpenseaAsset `json:"assets"`
}

// UnmarshalJSON will deserialize bytes into a OpenseaAssets model
func (openseaAssets *OpenseaAssets) UnmarshalJSON(d []byte) error {
	const (
		assetsJsonTag = "assets"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	if v := data.Value(assetsJsonTag); v != nil {
		for _, item := range data.Value(assetsJsonTag).([]interface{}) {
			bytes, _ := json.Marshal(item)
			obj := OpenseaAsset{}
			if err := json.Unmarshal(bytes, &obj); err != nil {
				return err
			}
			openseaAssets.ProtoAssets = append(openseaAssets.ProtoAssets, &obj)
		}
	}
	return nil
}
