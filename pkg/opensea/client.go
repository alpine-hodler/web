package opensea

import (
	"bytes"
	"net/http"

	"github.com/alpine-hodler/sdk/internal/client"
	"github.com/alpine-hodler/sdk/internal/log"
)

// C is the opensea client
type C struct {
	client.Parent
	client http.Client
}

func (opensea *C) setHeaders(hreq *http.Request, creq client.Request) (e error) {
	// TODO depricate getting key/passphrase/secret with secret keeper
	hreq.Header.Add("accept", "*/*")
	// hreq.Header.Add("content-type", "application/json")
	// hreq.Header.Add("User-Agent", "Go coinbase Pro Client 1.0")
	return
}

func (opensea *C) Do(creq client.Request) (*http.Response, error) {
	uri := "https://api.opensea.io" + creq.EndpointPath()
	client.Logf(log.DEBUG, &creq, `{Client:{URI:%s}}`, uri)
	hreq, err := http.NewRequest(creq.MethodStr(), uri, bytes.NewReader(creq.GetBody().Bytes()))
	if err != nil {
		return nil, err
	}
	if err := opensea.setHeaders(hreq, creq); err != nil {
		return nil, err
	}
	return opensea.client.Do(hreq)
}

// Connect creats a new client instance
func (opensea *C) Connect() error {
	opensea.client = http.Client{}
	return nil
}

// Identifier identifies requests
func (opensea *C) Identifier() string {
	return "Opensea"
}

func DefaultConnector() (client.C, error) {
	return &C{}, nil
}

// NewAccounts will return a new accounts structure to query on trading accounts
func NewClient(conn client.Connector) *C {
	openseaClient := new(C)
	client.ConstructParent(&openseaClient.Parent, conn)
	return openseaClient
}
