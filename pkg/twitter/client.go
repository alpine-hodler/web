package twitter

import (
	"bytes"
	"context"
	"net/http"

	"github.com/alpine-hodler/sdk/internal/client"
	"github.com/alpine-hodler/sdk/internal/env"
)

const (
	authorizationHeaderParam    = "Authorization"
	authorizationPrefix         = "OAuth"
	bearerHeaderPrefix          = "Bearer"
	contentType                 = "Content-Type"
	formContentType             = "application/x-www-form-urlencoded"
	defaultOauthSignatureMethod = "HMAC-SHA1"
	defaultOauthVersion         = "1.0"
	defaultTwitterAPIURL        = "https://api.twitter.com"
	identifier                  = "Twitter"
	oauthConsumerKeyParam       = "oauth_consumer_key"
	oauthNonceParam             = "oauth_nonce"
	oauthSignatureParam         = "oauth_signature"
	oauthSignatureMethodParam   = "oauth_signature_method"
	oauthTimestampParam         = "oauth_timestamp"
	oauthTokenParam             = "oauth_token"
	oauthVersionParam           = "oauth_version"

	TwitterISO8601 = "2006-01-02T15:04:05.000Z"
)

// C is the Twitter client.
type C struct {
	client.Parent
	client *http.Client
	nonce  string
	url    string
}

// NewClientAuth1 ...
func NewClientAuth1(_ context.Context) (*C, error) {
	c := new(C)
	c.client = new(http.Client)
	c.client.Transport = newTransportAuth1().
		setAccessToken(env.TwitterAccessToken.Get()).
		setAccessTokenSecret(env.TwitterAccessTokenSecret.Get()).
		setConsumerKey(env.TwitterAccessKey.Get()).
		setConsumerSecret(env.TwitterSecret.Get())

	c.setURL()
	client.ConstructParent(&c.Parent, func() (client.C, error) { return c, nil })
	return c, nil
}

// NewClientAuth2 will use an OAuth 2.0 Bearer Token.  This allows a Twitter developer app to access information
// publicly available on Twitter.  This will authenticates requests on behalf of your developer App. As this method is
// specific to the App, it does not involve any users. This method is typically for developers that need read-only
// access to public information.
func NewClientAuth2(_ context.Context) (*C, error) {
	c := new(C)
	c.setURL()

	c.client = new(http.Client)
	c.client.Transport = newTransportAuth2().setBearer(env.TwitterAuth2BearerToken.Get())

	client.ConstructParent(&c.Parent, func() (client.C, error) { return c, nil })
	return c, nil
}

// NewClientConnector will connect to the client using a custom connector.
func NewClientConnector(_ context.Context, conn client.Connector) (*C, error) {
	c := new(C)
	client.ConstructParent(&c.Parent, conn)
	return c, nil
}

// Do makes an http request to the Twitter API, given a method and an endpoint.
func (c *C) Do(creq client.Request) (*http.Response, error) {
	uri := c.url + creq.URIPostAuthority()
	hreq, err := http.NewRequest(creq.MethodStr(), uri, bytes.NewReader(creq.GetBody().Bytes()))
	if err != nil {
		return nil, err
	}
	return c.client.Do(hreq)
}

// Connect creats a new client instance
func (c *C) Connect() error { return nil }

// Identifier identifies requests
func (c *C) Identifier() string {
	return identifier
}

// setURL attempts to set the URL from the environment variables.  If there are none, it will set it to the default
// URL.
func (c *C) setURL() {
	if fromEnv := env.TwitterURL.Get(); len(fromEnv) > 0 {
		c.url = fromEnv
	} else {
		c.url = defaultTwitterAPIURL
	}
}
