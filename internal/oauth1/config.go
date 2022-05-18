package oauth1

import (
	"context"
	"net/http"
)

const (
	oauthTokenSecretParam       = "oauth_token_secret"
	oauthCallbackConfirmedParam = "oauth_callback_confirmed"
)

// NewClient returns a new http Client which signs requests via OAuth1.
func NewClient(ctx context.Context, consumerKey, consumerSecret, accessToken, accessTokenSecret string) *http.Client {
	transport := &Transport{
		// source: StaticTokenSource(token),
		consumerKey:       consumerKey,
		consumerSecret:    consumerSecret,
		accessToken:       accessToken,
		accessTokenSecret: accessTokenSecret,

		auther: newAuther(consumerKey, consumerSecret),
	}
	return &http.Client{Transport: transport}
}
