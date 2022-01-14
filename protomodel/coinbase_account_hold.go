package protomodel

import (
	"time"

	"github.com/alpine-hodler/sdk/serial"
)

// * This is a generated file, do not edit

// CoinbaseHold represents the hold on an account that belong to the same
// profile as the API key. Holds are placed on an account for any active orders
// or pending withdraw requests. As an order is filled, the hold amount is
// updated. If an order is canceled, any remaining hold is removed. For
// withdrawals, the hold is removed after it is completed.
type CoinbaseAccountHold struct {
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"id"`
	Ref       string    `json:"ref"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseAccountHold model
func (coinbaseAccountHold *CoinbaseAccountHold) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag        = "id"
		createdAtJsonTag = "created_at"
		updatedAtJsonTag = "updated_at"
		typeJsonTag      = "type"
		refJsonTag       = "ref"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(idJsonTag, &coinbaseAccountHold.Id)
	data.UnmarshalString(refJsonTag, &coinbaseAccountHold.Ref)
	data.UnmarshalString(typeJsonTag, &coinbaseAccountHold.Type)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &coinbaseAccountHold.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, updatedAtJsonTag, &coinbaseAccountHold.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
