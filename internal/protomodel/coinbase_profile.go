package protomodel

import (
	"time"

	"github.com/alpine-hodler/sdk/internal/serial"
)

// * This is a generated file, do not edit

// CoinbaseProfile represents a profile to interact with the API.
type CoinbaseProfile struct {
	Active    bool      `json:"active" bson:"active"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	HasMargin bool      `json:"has_margin" bson:"has_margin"`
	Id        string    `json:"id" bson:"id"`
	IsDefault bool      `json:"is_default" bson:"is_default"`
	Name      string    `json:"name" bson:"name"`
	UserId    string    `json:"user_id" bson:"user_id"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseProfile model
func (coinbaseProfile *CoinbaseProfile) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag        = "id"
		userIdJsonTag    = "user_id"
		nameJsonTag      = "name"
		activeJsonTag    = "active"
		isDefaultJsonTag = "is_default"
		hasMarginJsonTag = "has_margin"
		createdAtJsonTag = "created_at"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(activeJsonTag, &coinbaseProfile.Active)
	data.UnmarshalBool(hasMarginJsonTag, &coinbaseProfile.HasMargin)
	data.UnmarshalBool(isDefaultJsonTag, &coinbaseProfile.IsDefault)
	data.UnmarshalString(idJsonTag, &coinbaseProfile.Id)
	data.UnmarshalString(nameJsonTag, &coinbaseProfile.Name)
	data.UnmarshalString(userIdJsonTag, &coinbaseProfile.UserId)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &coinbaseProfile.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
