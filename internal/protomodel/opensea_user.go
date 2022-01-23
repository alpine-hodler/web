package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// The user of the an opensea asset owner
type OpenseaUser struct {
	Username string `json:"username" bson:"username"`
}

// UnmarshalJSON will deserialize bytes into a OpenseaUser model
func (openseaUser *OpenseaUser) UnmarshalJSON(d []byte) error {
	const (
		usernameJsonTag = "username"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(usernameJsonTag, &openseaUser.Username)
	return nil
}
