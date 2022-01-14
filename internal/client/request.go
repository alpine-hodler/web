package client

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alpine-hodler/sdk/pkg/model"

	"github.com/sirupsen/logrus"
)

type assignmentCallback func(interface{}, *Request) error

type Request struct {
	assignmentCallback assignmentCallback
	body               Body
	client             string
	connector          Connector
	endpoint           Endpoint
	endpointArgs       EndpointArgs
	errors             *Errors
	method             Method

	// slug is an 8 character randomly generated identifiery for the body, to be
	// used to identify request info in logging.
	slug string
}

func (req *Request) EndpointArgs() EndpointArgs { return req.endpointArgs }
func (req *Request) EndpointPath() string       { return req.endpoint.Path(req.endpointArgs) }
func (req *Request) Endpoint() Endpoint         { return req.endpoint }
func (req *Request) GetBody() *Body             { return &req.body }
func (req *Request) Method() Method             { return req.method }
func (req *Request) MethodStr() string          { return req.method.String() }

func New(conn Connector, method Method, endpoint Endpoint) *Request {
	req := new(Request)
	req.connector = conn
	req.endpoint = endpoint
	req.errors = new(Errors)
	req.method = method
	req.endpointArgs = make(EndpointArgs)
	return req
}

// generateSlub will make an 8 character randomly generated identifier for the body, which can be
// used to identify request info in logging.
func (req *Request) generateSlug() {
	b := make([]byte, 4)
	rand.Read(b)
	req.slug = hex.EncodeToString(b)
}

// parseErrorMessage takes a response and a status and builder an error message
// to send to the server
func parseErrorMessage(res *http.Response, status int) error {
	msg := model.ErrorMessage{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&msg); err != nil {
		return err
	}
	msg.Status = res.Status
	msg.StatusCode = http.StatusText(status)
	return fmt.Errorf("%+v", msg)
}

// validateResponse is a switch condition that parses an error response
func (req *Request) validateResponse(res *http.Response) (err error) {
	switch res.StatusCode {
	case http.StatusBadRequest:
		err = parseErrorMessage(res, http.StatusBadRequest)
	case http.StatusUnauthorized:
		err = parseErrorMessage(res, http.StatusUnauthorized)
	case http.StatusInternalServerError:
		err = parseErrorMessage(res, http.StatusInternalServerError)
	case http.StatusNotFound:
		err = parseErrorMessage(res, http.StatusNotFound)
	}
	if err != nil {
		logrus.Warn(Logf(req, err.Error()))
	}
	return
}

// AssignmentCallback will set a callback on the req that runs post-assignment
func (req *Request) AssignmentCallback(cb assignmentCallback) *Request {
	req.assignmentCallback = cb
	return req
}

// Fetch will use the req's connector to
func (req *Request) Fetch() *Assigner {
	assigner := newAssigner(req)
	// Generate the slug for identifying requests in async logging (if that ever
	// happens)
	req.generateSlug()

	// Then fetch the request from the API
	res, err := req.connector.fetch(req)
	if err != nil {
		req.errors.add(err)
	}

	// Validate the response, ensuring that there are no error codes or suspicious
	// messages
	req.errors.add(req.validateResponse(res))
	assigner.body = res.Body

	return assigner
}

func (req *Request) PathParam(key, value string) *Request {
	req.endpointArgs[key] = &EndpointArg{PathParam: &value}
	return req
}

// SetBody sets a body object on the request
func (req *Request) Body(body *Body) *Request {
	req.body = *body
	return req
}

func (req *Request) QueryParam(key string, value *string) *Request {
	req.endpointArgs[key] = &EndpointArg{QueryParam: func() (i *string) {
		if value != nil {
			i = value
		}
		return
	}()}
	return req
}
