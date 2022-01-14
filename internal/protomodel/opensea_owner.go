package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// The owner of an opensea asset
type OpenseaOwner struct {
	Address       string      `json:"address"`
	Config        string      `json:"config"`
	ProfileImgUrl string      `json:"profile_img_url"`
	ProtoUser     OpenseaUser `json:"user"`
}

// UnmarshalJSON will deserialize bytes into a OpenseaOwner model
func (openseaOwner *OpenseaOwner) UnmarshalJSON(d []byte) error {
	const (
		userJsonTag          = "user"
		profileImgUrlJsonTag = "profile_img_url"
		addressJsonTag       = "address"
		configJsonTag        = "config"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(addressJsonTag, &openseaOwner.Address)
	data.UnmarshalString(configJsonTag, &openseaOwner.Config)
	data.UnmarshalString(profileImgUrlJsonTag, &openseaOwner.ProfileImgUrl)
	openseaOwner.ProtoUser = OpenseaUser{}
	if err := data.UnmarshalStruct(userJsonTag, &openseaOwner.ProtoUser); err != nil {
		return err
	}
	return nil
}
