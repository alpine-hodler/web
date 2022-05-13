package client

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alpine-hodler/sdk/internal/log"
	"github.com/alpine-hodler/sdk/tools"
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

type ErrorMessage struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
}

// options is an interface for setting query parameters on a request.
type queryParamsOptions interface {
	SetQueryParams(*Request)
}

type bodyOptions interface {
	SetBody(*Request)
}

type stringer interface {
	String() string
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

// generateSlub will make an 8 character randomly generated identifier for the body, which can be used to identify
// request info in logging.
func (req *Request) generateSlug() {
	b := make([]byte, 4)
	rand.Read(b)
	req.slug = hex.EncodeToString(b)
}

// parseErrorMessage takes a response and a status and builder an error message to send to the server
func parseErrorMessage(res *http.Response, status int) error {
	msg := ErrorMessage{}
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
	if res == nil {
		err = fmt.Errorf("no response, check request and env file")
	} else {
		switch res.StatusCode {
		case http.StatusBadRequest:
			err = parseErrorMessage(res, http.StatusBadRequest)
		case http.StatusUnauthorized:
			err = parseErrorMessage(res, http.StatusUnauthorized)
		case http.StatusInternalServerError:
			err = parseErrorMessage(res, http.StatusInternalServerError)
		case http.StatusNotFound:
			err = parseErrorMessage(res, http.StatusNotFound)
		case http.StatusForbidden:
			err = parseErrorMessage(res, http.StatusForbidden)
		}
	}
	if err != nil {
		Logf(log.WARN, req, err.Error())
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
	// Generate the slug for identifying requests in async logging (if that ever happens)
	req.generateSlug()

	// Then fetch the request from the API
	res, err := req.connector.fetch(req)
	if err != nil {
		req.errors.add(err)
	}

	// Validate the response, ensuring that there are no error codes or suspicious messages
	if res != nil {
		req.errors.add(req.validateResponse(res))
		assigner.body = res.Body
	}

	return assigner
}

func (req *Request) PathParam(key, value string) *Request {
	req.endpointArgs[key] = &EndpointArg{PathParam: &value}
	return req
}

// SetBody sets a body object on the request
func (req *Request) SetBody(bodyType BodyType, opts interface{}) *Request {
	if opts != nil {
		b := NewBody(bodyType)
		req.body = *b
		opts.(bodyOptions).SetBody(req)
	}
	return req
}

// TODO
func (req *Request) SetBodyBool(key string, val *bool) *Request {
	if val != nil {
		req.body.data[key] = *val
	}
	return req
}

// TODO
func (req *Request) SetBodyFloat(key string, val *float64) *Request {
	if val != nil {
		req.body.data[key] = *val
	}
	return req
}

// TODO
func (req *Request) SetBodyInt(key string, val *int) *Request {
	if val != nil {
		req.body.data[key] = *val
	}
	return req
}

// TODO
func (req *Request) SetBodyString(key string, val *string) *Request {
	if val != nil {
		req.body.data[key] = *val
	}
	return req
}

// TODO
func (req *Request) SetStringer(key string, val stringer) *Request {
	if val != nil {
		req.body.data[key] = val.String()
	}
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

// SetQueryParamBool will set a query param from a bool value.
func (req *Request) SetQueryParamBool(key string, value *bool) *Request {
	req.endpointArgs[key] = &EndpointArg{QueryParam: func() (i *string) {
		if value != nil {
			boolStr := strconv.FormatBool(*value)
			i = &boolStr
		}
		return
	}()}
	return req
}

// SetQueryParamInt32 will set a query param from an int32 value.
func (req *Request) SetQueryParamInt32(key string, value *int32) *Request {
	req.endpointArgs[key] = &EndpointArg{QueryParam: func() (i *string) {
		if value != nil {
			i = tools.Int32PtrStringPtr(value)
		}
		return
	}()}
	return req
}

// SetQueryParamInt will set a query param from an int value.
func (req *Request) SetQueryParamInt(key string, value *int) *Request {
	req.endpointArgs[key] = &EndpointArg{QueryParam: func() (i *string) {
		if value != nil {
			i = tools.IntPtrStringPtr(value)
		}
		return
	}()}
	return req
}

// SetQueryParamString will set a query param from a string value.
func (req *Request) SetQueryParamString(key string, value *string) *Request {
	req.endpointArgs[key] = &EndpointArg{QueryParam: func() (i *string) {
		if value != nil {
			i = value
		}
		return
	}()}
	return req
}

// SetQueryParamStrings will set potentially many query param from a string slice.
func (req *Request) SetQueryParamStrings(key string, values []string) *Request {
	req.endpointArgs[key] = &EndpointArg{QueryParam: func() (i *string) {
		if values != nil {
			slice := []string{}
			for _, v := range values {
				slice = append(slice, v)
			}
			tmp := strings.Join(slice, ", ")
			i = &tmp
		}
		return
	}()}
	return req
}

// SetQueryParamTime will set a query param from a time.Time value.
func (req *Request) SetQueryParamTime(key string, value *time.Time) *Request {
	req.endpointArgs[key] = &EndpointArg{QueryParam: func() (i *string) {
		if value != nil {
			tmp := value.String()
			i = &tmp
		}
		return
	}()}
	return req
}

// SetQueryParams will take an opts interface and if it isn't nil, it will type cast it as an Options interface{} and
// attemp to set query parameters on the request.
func (req *Request) SetQueryParams(opts interface{}) *Request {
	if opts != nil {
		opts.(queryParamsOptions).SetQueryParams(req)
	}
	return req
}
