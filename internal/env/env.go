package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Variable int

const (
	_ Variable = iota
	AlpineHodlerLogLevel
	CoinbaseProURL
	CoinbaseProAccessPassphrase
	CoinbaseProAccessKey
	CoinbaseProSecret
	IEXURL
	IEXKey
	IEXSecret
	KrakenURL
	KrakenKey
	KrakenSecret
	TwitterAuth2BearerToken
	TwitterAccessToken
	TwitterAccessTokenSecret
	TwitterAccessKey
	TwitterSecret
	TwitterURL
)

const (
	alpineHodlerLogLevel          = "ALPINE_HODLER_LOG_LEVEL"
	coinbaseProURLID              = "CB_PRO_URL"
	coinbaseProAccessPassphraseID = "CB_PRO_ACCESS_PASSPHRASE"
	coinbaseProAccessKeyID        = "CB_PRO_ACCESS_KEY"
	coinbaseProSecretID           = "CB_PRO_SECRET"
	twitterAuth2BearerToken       = "TWITTER_BEARER_TOKEN"
	twitterAccessToken            = "TWITTER_ACCESS_TOKEN"
	twitterAccessTokenSecret      = "TWITTER_ACCESS_TOKEN_SECRET"
	twitterAccessKey              = "TWITTER_ACCESS_KEY"
	twitterSecret                 = "TWITTER_SECRET"
	twitterURL                    = "TWITTER_URL"
)

func Load(filepath string) {
	if err := godotenv.Load(filepath); err != nil {
		panic(fmt.Errorf("Error loading .env file: %v", err))
	}
	return
}

// Name will return the variable name of the env variable.
func (variable Variable) String() string {
	return map[Variable]string{
		AlpineHodlerLogLevel:        alpineHodlerLogLevel,
		CoinbaseProURL:              coinbaseProURLID,
		CoinbaseProAccessPassphrase: coinbaseProAccessPassphraseID,
		CoinbaseProAccessKey:        coinbaseProAccessKeyID,
		CoinbaseProSecret:           coinbaseProSecretID,
		TwitterAuth2BearerToken:     twitterAuth2BearerToken,
		TwitterAccessToken:          twitterAccessToken,
		TwitterAccessTokenSecret:    twitterAccessTokenSecret,
		TwitterAccessKey:            twitterAccessKey,
		TwitterSecret:               twitterSecret,
		TwitterURL:                  twitterURL,
	}[variable]
}

// Get will return the environment variable for a variable-type constant as a string.
func (variable Variable) Get() string {
	return os.Getenv(variable.String())
}

// SetAlpineHodlerLogLevel will set the ALPINE_HODLER_LOG_LEVEL environment variable
var SetAlpineHodlerLogLevel = func(level string) error {
	return os.Setenv(alpineHodlerLogLevel, level)
}

// SetCoinbaseProURL will set the CB_PRO_URL environment variable
var SetCoinbaseProURL = func(url string) error {
	return os.Setenv(coinbaseProURLID, url)
}

// SetCoinbaseProAccessPassphrase will set the CB_PRO_ACCESS_PASSPHRASE environment variable
var SetCoinbaseProAccessPassphrase = func(url string) error {
	return os.Setenv(coinbaseProAccessPassphraseID, url)
}

// SetCoinbaseProAccessKey will set the CB_PRO_ACCESS_KEY environment variable
var SetCoinbaseProAccessKey = func(url string) error {
	return os.Setenv(coinbaseProAccessKeyID, url)
}

// SetCoinbaseProSecret will set the CB_PRO_SECRET environment variable
var SetCoinbaseProSecret = func(url string) error {
	return os.Setenv(coinbaseProSecretID, url)
}

// SetTwitterAuth2BearerToken will set the TWITTER_AUTH2_BEARER_TOKEN environment variable
var SetTwitterAuth2BearerToken = func(token string) error {
	return os.Setenv(twitterAuth2BearerToken, token)
}

// SetTwitterURL will set the TWITTER_URL environment variable
var SetTwitterURL = func(url string) error {
	return os.Setenv(twitterURL, url)
}
