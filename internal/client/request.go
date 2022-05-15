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
	errors             *Errors
	method             Method
	uriBuilderAccesor  tools.URIBuilderAccessor
	uriBuilder         tools.URIBuilder

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

func (req *Request) URIBuilder() tools.URIBuilder                 { return req.uriBuilder }
func (req *Request) URIBuilderAccessor() tools.URIBuilderAccessor { return req.uriBuilderAccesor }
func (req *Request) GetBody() *Body                               { return &req.body }
func (req *Request) Method() Method                               { return req.method }
func (req *Request) MethodStr() string                            { return req.method.String() }

func New(conn Connector, method Method, accessor tools.URIBuilderAccessor) *Request {
	req := new(Request)
	req.connector = conn
	req.errors = new(Errors)
	req.method = method
	req.uriBuilder = make(tools.URIBuilder)
	req.uriBuilderAccesor = accessor
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

// AddPathParam will add a path parameter to the tools.URIBuilder.
func (req *Request) SetPathParam(key, value string) *Request {
	req.uriBuilder.Add(tools.URIBuilderComponentPath, key, value)
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

// SetQueryParamBool will set a query param from a bool value.
func (req *Request) SetQueryParamBool(key string, value *bool) *Request {
	if value != nil {
		req.uriBuilder.Add(tools.URIBuilderComponentQuery, key, strconv.FormatBool(*value))
	}
	return req
}

// SetQueryParamInt32 will set a query param from an int32 value.
func (req *Request) SetQueryParamInt32(key string, value *int32) *Request {
	if value != nil {
		req.uriBuilder.Add(tools.URIBuilderComponentQuery, key, tools.Int32PtrString(value))
	}
	return req
}

// SetQueryParamInt will set a query param from an int value.
func (req *Request) SetQueryParamInt(key string, value *int) *Request {
	if value != nil {
		req.uriBuilder.Add(tools.URIBuilderComponentQuery, key, tools.IntPtrString(value))
	}
	return req
}

// SetQueryParamString will set a query param from a string value.
func (req *Request) SetQueryParamString(key string, value *string) *Request {
	if value != nil {
		req.uriBuilder.Add(tools.URIBuilderComponentQuery, key, *value)
	}
	return req
}

// SetQueryParamStrings will set potentially many query param from a string slice.
func (req *Request) SetQueryParamStrings(key string, values []string) *Request {
	if values != nil {
		// TODO: Add this logic to the tools package
		slice := []string{}
		for _, v := range values {
			slice = append(slice, v)
		}
		req.uriBuilder.Add(tools.URIBuilderComponentQuery, key, strings.Join(slice, ", "))
	}
	return req
}

// SetQueryParamTime will set a query param from a time.Time value.
func (req *Request) SetQueryParamTime(key string, value *time.Time) *Request {
	if value != nil {
		req.uriBuilder.Add(tools.URIBuilderComponentQuery, key, value.String())
	}
	return req
}

// SetQueryParams will take an opts interface and if it isn't nil, it will type cast it as an Options interface{} and
// attempt to set query parameters on the request.
func (req *Request) SetQueryParams(opts interface{}) *Request {
	if opts != nil {
		opts.(queryParamsOptions).SetQueryParams(req)
	}
	return req
}

// URIPostAuthority will return everything after the authority portion of the URI
func (req *Request) URIPostAuthority() string {
	return req.uriBuilderAccesor.PostAuthority(req.uriBuilder)
}
