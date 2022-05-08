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
)

const (
	alpineHodlerLogLevel          = "ALPINE_HODLER_LOG_LEVEL"
	coinbaseProURLID              = "CB_PRO_URL"
	coinbaseProAccessPassphraseID = "CB_PRO_ACCESS_PASSPHRASE"
	coinbaseProAccessKeyID        = "CB_PRO_ACCESS_KEY"
	coinbaseProSecretID           = "CB_PRO_SECRET"
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
	}[variable]
}

// Get will return the environment variable for a variable-type constant as a string.
func (variable Variable) Get() string {
	return os.Getenv(variable.String())
}

// SetCoinbaseProURL will set the CB_PRO_URL environment variable
func SetCoinbaseProURL(url string) error {
	return os.Setenv(coinbaseProURLID, url)
}

// SetCoinbaseProAccessPassphrase will set the CB_PRO_ACCESS_PASSPHRASE environment variable
func SetCoinbaseProAccessPassphrase(url string) error {
	return os.Setenv(coinbaseProAccessPassphraseID, url)
}

// SetCoinbaseProAccessKey will set the CB_PRO_ACCESS_KEY environment variable
func SetCoinbaseProAccessKey(url string) error {
	return os.Setenv(coinbaseProAccessKeyID, url)
}

// SetCoinbaseProSecret will set the CB_PRO_SECRET environemnt variable
func SetCoinbaseProSecret(url string) error {
	return os.Setenv(coinbaseProSecretID, url)
}
