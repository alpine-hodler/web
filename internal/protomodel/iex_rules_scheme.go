package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// IEXRulesScheme is one of the latest schemes for data points, notification
// types, and operators used to construct rules.
type IexRulesScheme struct {
	IsLookup  bool    `json:"isLookup"`
	Label     string  `json:"label"`
	Scope     string  `json:"scope"`
	Type      string  `json:"type"`
	Value     string  `json:"value"`
	Weight    float64 `json:"weight"`
	WeightKey string  `json:"weightKey"`
}

// UnmarshalJSON will deserialize bytes into a IexRulesScheme model
func (iexRulesScheme *IexRulesScheme) UnmarshalJSON(d []byte) error {
	const (
		labelJsonTag     = "label"
		valueJsonTag     = "value"
		typeJsonTag      = "type"
		scopeJsonTag     = "scope"
		isLookupJsonTag  = "isLookup"
		weightJsonTag    = "weight"
		weightKeyJsonTag = "weightKey"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(isLookupJsonTag, &iexRulesScheme.IsLookup)
	data.UnmarshalFloat(weightJsonTag, &iexRulesScheme.Weight)
	data.UnmarshalString(labelJsonTag, &iexRulesScheme.Label)
	data.UnmarshalString(scopeJsonTag, &iexRulesScheme.Scope)
	data.UnmarshalString(typeJsonTag, &iexRulesScheme.Type)
	data.UnmarshalString(valueJsonTag, &iexRulesScheme.Value)
	data.UnmarshalString(weightKeyJsonTag, &iexRulesScheme.WeightKey)
	return nil
}
