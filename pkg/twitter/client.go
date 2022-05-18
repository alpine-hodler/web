package twitter

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/alpine-hodler/sdk/internal/client"
	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/internal/oauth1"
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

const (
	TwitterISO8601 = "2006-01-02T15:04:05.000Z"
)

// C is the Twitter client
type C struct {
	client.Parent

	client *http.Client

	authenticationMethod AuthenticationMethod

	nonce string

	accessKey         string
	secret            string
	accessTokenSecret string
	accessToken       string

	auth2BearerToken string
	url              string
}

func NewClientAuth1(_ context.Context) (*C, error) {
	c := new(C)

	c.authenticationMethod = Auth1aUserContext
	c.accessKey = env.TwitterAccessKey.Get()
	c.secret = env.TwitterSecret.Get()
	c.accessTokenSecret = env.TwitterAccessTokenSecret.Get()
	c.accessToken = env.TwitterAccessToken.Get()

	c.url = env.TwitterURL.Get()
	client.ConstructParent(&c.Parent, func() (client.C, error) { return c, nil })
	return c, nil
}

// NewClientAuth2 will use an OAuth 2.0 Bearer Token.  This allows a Twitter developer app to access information
// publicly available on Twitter.  This will authenticates requests on behalf of your developer App. As this method is
// specific to the App, it does not involve any users. This method is typically for developers that need read-only
// access to public information.
func NewClientAuth2(_ context.Context) (*C, error) {
	c := new(C)
	c.auth2BearerToken = env.TwitterAuth2BearerToken.Get()
	c.authenticationMethod = Auth2BearerToken
	c.url = env.TwitterURL.Get()
	client.ConstructParent(&c.Parent, func() (client.C, error) { return c, nil })
	return c, nil
}

// NewClientConnector will connect to the client using a custom connector.
func NewClientConnector(_ context.Context, conn client.Connector) (*C, error) {
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

func (c *C) doDefault(creq client.Request) (*http.Response, error) {
	// uri := c.url + creq.URIPostAuthority()
	uri := "https://api.twitter.com/1.1/statuses/home_timeline.json?count=2"
	fmt.Println(creq.Method())
	hreq, err := http.NewRequest(creq.MethodStr(), uri, bytes.NewReader(creq.GetBody().Bytes()))
	if err != nil {
		return nil, err
	}
	if err := c.setHeaders(hreq, creq); err != nil {
		return nil, err
	}
	fmt.Printf("-->%+v\n", c.client)
	return c.client.Do(hreq)
}

// Do  makes an http request to the Twitter API, given a method and an endpoint.
func (c *C) Do(creq client.Request) (*http.Response, error) {
	return c.doDefault(creq)
}

// Connect creats a new client instance
func (c *C) Connect() error {
	switch c.authenticationMethod {
	case Auth1aUserContext:
		// config := oauth1.NewConfig(c.accessKey, c.secret)
		// token := oauth1.NewToken(c.accessToken, c.accessTokenSecret)
		c.client = oauth1.NewClient(context.TODO(), c.accessKey, c.secret, c.accessToken, c.accessTokenSecret)

		// fmt.Println("token:", token)

		fmt.Printf("->%+v\n", c.client)
	default:
		c.client = &http.Client{}
	}
	return nil
}

// Identifier identifies requests
func (c *C) Identifier() string {
	return "Twitter"
}

// &{1334223508788932611-NVoJFJxVr89p2quJ91B8sE4BPe2Igc LJfRG8RfQXl9L4RQt9hQgJK4D75gy7F9kyNKSL2FC00T5glsD5}
// &{1334223508788932611-NVoJFJxVr89p2quJ91B8sE4BPe2Igc y06ItlOlMryicPQYt6Mw513ySZyjOUiX99SsK6VYcY9mm}

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
