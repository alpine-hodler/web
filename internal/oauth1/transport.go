package oauth1

import (
	"fmt"
	"net/http"
)

// Transport is an http.RoundTripper which makes OAuth1 HTTP requests. It
// wraps a base RoundTripper and adds an Authorization header using the
// token from a TokenSource.
//
// Transport is a low-level component, most users should use Config to create
// an http.Client instead.
type Transport struct {
	consumerKey       string
	consumerSecret    string
	accessTokenSecret string
	accessToken       string

	// auther adds OAuth1 Authorization headers to requests
	auther *auther
}

// RoundTrip authorizes the request with a signed OAuth1 Authorization header
// using the auther and TokenSource.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// if t.source == nil {
	// 	return nil, fmt.Errorf("oauth1: Transport's source is nil")
	// }
	// accessToken, err := t.source.Token()
	// if err != nil {
	// 	return nil, err
	// }
	if t.auther == nil {
		return nil, fmt.Errorf("oauth1: Transport's auther is nil")
	}
	// RoundTripper should not modify the given request, clone it
	// req2 := cloneRequest(req)
	err := t.auther.setRequestAuthHeader(req, t.accessToken, t.accessTokenSecret)
	if err != nil {
		return nil, err
	}
	return http.DefaultTransport.RoundTrip(req)
}
