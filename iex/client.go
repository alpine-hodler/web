package iex

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/alpine-hodler/sdk/env"

	"github.com/alpine-hodler/sdk/client"

	"github.com/sirupsen/logrus"
)

// C is the iex client
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

func authorizedURI(creq client.Request) string {
	unauthorizedURI := env.IEXURL.Get() + creq.EndpointPath()
	u, _ := url.Parse(unauthorizedURI)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("token", env.IEXKey.Get())
	u.RawQuery = q.Encode()
	return u.String()
}

func (iex *C) Do(creq client.Request) (*http.Response, error) {
	uri := authorizedURI(creq)
	logrus.Debug(client.Logf(&creq, `{Client:{URI:%s}}`, uri))
	hreq, err := http.NewRequest(creq.MethodStr(), uri, bytes.NewReader(creq.GetBody().Bytes()))
	if err != nil {
		return nil, err
	}
	return iex.client.Do(hreq)
}

// Connect creats a new client instance
func (iex *C) Connect() error {
	iex.client = http.Client{}
	return nil
}

// Identifier identifies requests
func (iex *C) Identifier() string {
	return "IEX"
}

func DefaultClient() (client.C, error) {
	return &C{}, nil
}

func NewClient(conn client.Connector) *C {
	iexClient := new(C)
	client.ConstructParent(&iexClient.Parent, conn)
	return iexClient
}
