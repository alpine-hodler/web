package kraken

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/alpine-hodler/sdk/internal/env"

	"github.com/alpine-hodler/sdk/internal/client"

	"github.com/sirupsen/logrus"
)

// C is the kraken client
type C struct {
	client.Parent
	client http.Client
}

func (cb *C) generateSig(path string, values url.Values, secret []byte) string {
	sha := sha256.New()
	sha.Write([]byte(values.Get("nonce") + values.Encode()))
	shasum := sha.Sum(nil)

	mac := hmac.New(sha512.New, secret)
	mac.Write(append([]byte(path), shasum...))
	macsum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(macsum)
}

func (kraken *C) setHeaders(hreq *http.Request, creq client.Request) (e error) {
	// TODO depricate getting key/passphrase/secret with secret keeper
	var (
		payload             = url.Values{}
		apiSecret           = env.KrakenSecret.Get()
		b64DecodedSecret, _ = base64.StdEncoding.DecodeString(apiSecret)
		sig                 = kraken.generateSig(creq.EndpointPath(), payload, b64DecodedSecret)
	)
	hreq.Header.Add("accept", "application/json")
	hreq.Header.Add("content-type", "application/json")
	// hreq.Header.Add("User-Agent", "Go coinbase Pro Client 1.0")
	hreq.Header.Add("API-Sign", sig)

	logMsg := `{Client:{Access:{API-Sign:%s}}}`
	logrus.Debug(client.Logf(&creq, logMsg, sig))
	return
}

func (kraken *C) Do(creq client.Request) (*http.Response, error) {
	uri := env.KrakenURL.Get() + creq.EndpointPath()
	logrus.Debug(client.Logf(&creq, `{Client:{URI:%s}}`, uri))
	hreq, err := http.NewRequest(creq.MethodStr(), uri, bytes.NewReader(creq.GetBody().Bytes()))
	if err != nil {
		return nil, err
	}
	if err := kraken.setHeaders(hreq, creq); err != nil {
		return nil, err
	}
	return kraken.client.Do(hreq)
}

// Connect creats a new client instance
func (kraken *C) Connect() error {
	kraken.client = http.Client{}
	return nil
}

// Identifier identifies requests
func (kraken *C) Identifier() string {
	return "Kraken"
}

func DefaultConnector() (client.C, error) {
	return &C{}, nil
}

func NewClient(conn client.Connector) *C {
	krakenClient := new(C)
	client.ConstructParent(&krakenClient.Parent, conn)
	return krakenClient
}
