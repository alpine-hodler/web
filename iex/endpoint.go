package iex

import (
	"path"

	"github.com/alpine-hodler/sdk/client"
)

// * This is a generated file, do not edit
type Endpoint int

const (
	_ Endpoint = iota
	RulesEndpoint
	RulesSchemaEndpoint
)

// Rules for automated alert
func RulesPath(args client.EndpointArgs) string {
	return path.Join("/stable", "rules", "lookup", *args["value"].PathParam)
}

// Pull the latest schema for data points, notification types, and operators
// used to construct rules. If a schema object has "isLookup": true, pass the
// value key to /stable/rules/lookup/{value}. This returns all valid values for
// the rightValue of a condition.
func RulesSchemaPath(args client.EndpointArgs) string {
	return path.Join("/stable", "rules", "schema")
}

// Get takes an endpoint const and endpoint arguments to parse the URL endpoint
// path.
func (endpoint Endpoint) Path(args client.EndpointArgs) string {
	return map[Endpoint]func(args client.EndpointArgs) string{
		RulesEndpoint:       RulesPath,
		RulesSchemaEndpoint: RulesSchemaPath,
	}[endpoint](args)
}
