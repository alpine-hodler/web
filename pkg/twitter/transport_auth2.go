package twitter

import (
	"fmt"
	"net/http"
)

// TODO
type transportAuth2 struct {
	bearer string
}

// NewTransportAuth1 will return a transportAuth2
func newTransportAuth2() *transportAuth2 {
	return new(transportAuth2)
}

// setBearer will set the bearer field on transportAuth1.
func (rt *transportAuth2) setBearer(val string) *transportAuth2 {
	rt.bearer = val
	return rt
}

// RoundTrip authorizes the request with a signed OAuth1 Authorization header
// using the auther and TokenSource.
func (rt *transportAuth2) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(authorizationHeaderParam, fmt.Sprintf("%s %s", bearerHeaderPrefix, rt.bearer))
	return http.DefaultTransport.RoundTrip(req)
}
