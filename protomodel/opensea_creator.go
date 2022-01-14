package protomodel

import "github.com/alpine-hodler/sdk/serial"

// * This is a generated file, do not edit

// The creator of the an opensea asset
type OpenseaCreator struct {
	Address       string      `json:"address"`
	Config        string      `json:"config"`
	ProfileImgUrl string      `json:"profile_img_url"`
	ProtoUser     OpenseaUser `json:"user"`
}

// UnmarshalJSON will deserialize bytes into a OpenseaCreator model
func (openseaCreator *OpenseaCreator) UnmarshalJSON(d []byte) error {
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
	data.UnmarshalString(addressJsonTag, &openseaCreator.Address)
	data.UnmarshalString(configJsonTag, &openseaCreator.Config)
	data.UnmarshalString(profileImgUrlJsonTag, &openseaCreator.ProfileImgUrl)
	openseaCreator.ProtoUser = OpenseaUser{}
	if err := data.UnmarshalStruct(userJsonTag, &openseaCreator.ProtoUser); err != nil {
		return err
	}
	return nil
}
