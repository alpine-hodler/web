package kraken

import (
	"path"

	"github.com/alpine-hodler/sdk/internal/client"
)

// * This is a generated file, do not edit
type Endpoint int

const (
	_ Endpoint = iota
	ServerTimeEndpoint
	SystemStatusEndpoint
)

// Get the server's time.
func ServerTimePath(args client.EndpointArgs) string {
	return path.Join("/0", "public", "Time")
}

// Get the current system status or trading mode.
func SystemStatusPath(args client.EndpointArgs) string {
	return path.Join("/0", "public", "SystemStatus")
}

// Get takes an endpoint const and endpoint arguments to parse the URL endpoint path.
func (endpoint Endpoint) Path(args client.EndpointArgs) string {
	return map[Endpoint]func(args client.EndpointArgs) string{
		ServerTimeEndpoint:   ServerTimePath,
		SystemStatusEndpoint: SystemStatusPath,
	}[endpoint](args)
}
