package client

import (
	"net/http"
)

// C is the client interface
//go:generate mockgen -source client.go -destination ../client/client_mock.go -package client
type C interface {
	Identifier() string
	Connect() error
	Do(Request) (*http.Response, error)
}

// connector is a client connector, such as the New function.  It's broken into
// it's own type for decoding purposes.
type Connector func() (C, error)

func (conn Connector) fetch(req *Request) (*http.Response, error) {
	c, err := conn()
	if err != nil {
		return nil, err
	}
	c.Connect()
	req.client = c.Identifier()
	LogHTTPRequest(req)
	// req.logHTTPRequest()
	// req.logHTTPRequestBody()

	var res *http.Response
	res, err = c.Do(*req)

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Raw Response Body:\n%v\n", string(body))

	if err != nil {
		return nil, err
	}

	LogResponseStatus(req, res)
	return res, nil
}
