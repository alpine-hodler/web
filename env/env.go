package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Variable int

const (
	_ Variable = iota
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

func Load(filepath string) {
	if err := godotenv.Load(filepath); err != nil {
		panic(fmt.Errorf("Error loading .env file: %v", err))
	}
	return
}

// Name will return the variable name of the env variable.
func (variable Variable) Name() string {
	return map[Variable]string{
		CoinbaseProURL:              "CB_PRO_URL",
		CoinbaseProAccessPassphrase: "CB_PRO_ACCESS_PASSPHRASE",
		CoinbaseProAccessKey:        "CB_PRO_ACCESS_KEY",
		CoinbaseProSecret:           "CB_PRO_SECRET",
		IEXURL:                      "IEX_URL",
		IEXKey:                      "IEX_KEY",
		IEXSecret:                   "IEX_SECRET",
		KrakenURL:                   "KRAKEN_URL",
		KrakenKey:                   "KRAKEN_KEY",
		KrakenSecret:                "KRAKEN_SECRET",
	}[variable]
}

// Get will return the environment variable for a variable-type constant as a
// string.
func (variable Variable) Get() string {
	return os.Getenv(variable.Name())
}
