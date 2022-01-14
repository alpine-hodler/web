package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// The display type for a collection
type OpenseaDisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}

// UnmarshalJSON will deserialize bytes into a OpenseaDisplayData model
func (openseaDisplayData *OpenseaDisplayData) UnmarshalJSON(d []byte) error {
	const (
		cardDisplayStyleJsonTag = "card_display_style"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(cardDisplayStyleJsonTag, &openseaDisplayData.CardDisplayStyle)
	return nil
}
