package client

import (
	"net/http"
)

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

	//code snippet
	// bodyBytes, err := ioutil.ReadAll(res.Body)
	//print result
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)

	// fmt.Println(res)
	if err != nil {
		return nil, err
	}

	LogResponseStatus(req, res)
	return res, nil
}
