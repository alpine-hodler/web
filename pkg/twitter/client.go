package twitter

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/alpine-hodler/sdk/internal/client"
	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/internal/log"
)

// AuthenciationMethod types the different kinds of authentication method twitter offers
//
// source: https://developer.twitter.com/en/docs/authentication/overview
type AuthenticationMethod uint8

const (
	// OAuth 1.0a allows an authorized Twitter developer App to access private account information or perform a Twitter
	// action on behalf of a Twitter account.
	Auth1aUserContext AuthenticationMethod = iota

	// OAuth 2.0 Bearer Token allows a Twitter developer app to access information publicly available on Twitter.  This
	// will authenticates requests on behalf of your developer App. As this method is specific to the App, it does not
	// involve any users. This method is typically for developers that need read-only access to public information.
	Auth2BearerToken

	// OAuth 2.0 User Context allows you to authenticate on behalf of another account with greater control over an
	// application’s scope, and authorization flows across multiple devices.
	Auth2PKCE

	// Many of Twitter’s enterprise APIs require the use of HTTP Basic Authentication.
	BasicAuthentication
)

// C is the Twitter client
type C struct {
	client.Parent

	client http.Client

	authenticationMethod AuthenticationMethod
	auth2BearerToken     string
	url                  string
}

// NewAuth2Client will use an OAuth 2.0 Bearer Token.  This allows a Twitter developer app to access information
// publicly available on Twitter.  This will authenticates requests on behalf of your developer App. As this method is
// specific to the App, it does not involve any users. This method is typically for developers that need read-only
// access to public information.
func NewAuth2Client() (*C, error) {
	c := new(C)
	c.auth2BearerToken = env.TwitterAuth2BearerToken.Get()
	c.authenticationMethod = Auth2BearerToken
	c.url = env.TwitterURL.Get()
	client.ConstructParent(&c.Parent, func() (client.C, error) { return c, nil })
	return c, nil
}

// NewClientConnector will connect to the client using a custom connector.
func NewClientConnector(conn client.Connector) (*C, error) {
	c := new(C)
	client.ConstructParent(&c.Parent, conn)
	return c, nil
}

// setHeaders sets the headers for a twitter API request.
func (c *C) setHeaders(hreq *http.Request, creq client.Request) (e error) {
	switch c.authenticationMethod {
	case Auth2BearerToken:
		bearer := fmt.Sprintf("bearer %s", c.auth2BearerToken)
		hreq.Header.Add("authorization", bearer)
	}
	return
}

// Do  makes an http request to the Twitter API, given a method and an endpoint.
func (c *C) Do(creq client.Request) (*http.Response, error) {
	uri := c.url + creq.URIPostAuthority()

	client.Logf(log.DEBUG, &creq, `{Client:{URI:%s}}`, uri)

	hreq, err := http.NewRequest(creq.MethodStr(), uri, bytes.NewReader(creq.GetBody().Bytes()))
	if err != nil {
		return nil, err
	}
	if err := c.setHeaders(hreq, creq); err != nil {
		return nil, err
	}
	return c.client.Do(hreq)
}

// Connect creats a new client instance
func (c *C) Connect() error {
	c.client = http.Client{}
	return nil
}

// Identifier identifies requests
func (c *C) Identifier() string {
	return "Twitter"
}

// SetAuthenticationMethod will set an AuthenticationMethod on the client.
func (c *C) SetAuthenticationMethod(method AuthenticationMethod) *C {
	c.authenticationMethod = method
	return c
}

// SetAuth2BearerToken will set an Oauth 2.0 bearer token on the client.
func (c *C) SetAuth2BearerToken(token string) *C {
	c.auth2BearerToken = token
	return c
}

// SetURL will set the URL to create HTTP requests against on the client.
func (c *C) SetURL(url string) *C {
	c.url = url
	return c
}
