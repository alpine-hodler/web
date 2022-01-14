package opensea

import (
	"path"
	"strings"

	"github.com/alpine-hodler/sdk/client"
)

// * This is a generated file, do not edit
type Endpoint int

const (
	_ Endpoint = iota
	AssetsEndpoint
)

// Get alll assets
func AssetsPath(args client.EndpointArgs) (p string) {
	p = path.Join("/api", "v1", "assets")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Get takes an endpoint const and endpoint arguments to parse the URL endpoint
// path.
func (endpoint Endpoint) Path(args client.EndpointArgs) string {
	return map[Endpoint]func(args client.EndpointArgs) string{
		AssetsEndpoint: AssetsPath,
	}[endpoint](args)
}
