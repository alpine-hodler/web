package client

import (
	"net/url"
	"strings"
)

// EndpointArg is a wrapper to parse parameters for endpoints
type EndpointArg struct {
	PathParam  *string
	QueryParam *string
}

type EndpointArgs map[string]*EndpointArg

// Endpoint is an interface to build an endpiont from a list of string args
type Endpoint interface {
	// Path takes a list of strings as arguments to build an endpoint
	Path(EndpointArgs) string
}

// QueryPath get the query path for an endpoint arg
func (args EndpointArgs) QueryPath() *url.URL {
	u, _ := url.Parse("")
	q, _ := url.ParseQuery(u.RawQuery)
	for key, arg := range args {
		if param := arg.QueryParam; param != nil {
			q.Add(key, *param)
		}
	}
	u.RawQuery = q.Encode()
	return u
}

func JoinEndpointParts(parts ...string) string {
	return "/" + strings.Join(parts, "/")
}
