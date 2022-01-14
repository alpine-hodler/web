package iex

import (
	"github.com/alpine-hodler/sdk/pkg/model"
)

// Rule for automated alerts
//
// * source: https://iexcloud.io/docs/api/#lookup-values
func (client *C) Rules(value string) (m []*model.IexRule, err error) {
	req := client.Get(RulesEndpoint)
	return m, req.PathParam("value", value).Fetch().Assign(&m).JoinMessages()
}

// Pull the latest schema for data points, notification types, and operators
// used to construct rules.
//
// * source: https://iexcloud.io/docs/api/#rules-schema
func (client *C) RulesSchema() (m *model.IexRulesSchema, err error) {
	req := client.Get(RulesSchemaEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}
