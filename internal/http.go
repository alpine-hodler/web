package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Options are optional data than can be encoded into a request's URL or Body.
type Options interface {
	EncodeBody() (io.Reader, error)
	EncodeQuery(*http.Request)
}

type endpoint interface {
	Path(map[string]string) string

	// Scope is the OAuth 2.0 scope required to make requests to this endpoint.
	Scope() string
}

// stringer is any type that can textualize it's value as a string.
type stringer interface {
	String() string
}

func httpQueryEncode(req *http.Request, key, val string) {
	q := req.URL.Query()
	q.Add(key, val)
	req.URL.RawQuery = q.Encode()
}

// HTTPQueryEncodeBool will convert a bool pointer into a string and then add it to the query parameters of an HTTP
// request's URL.
func HTTPQueryEncodeBool(req *http.Request, key string, val *bool) {
	if val != nil {
		httpQueryEncode(req, key, strconv.FormatBool(*val))
	}
}

func intPtrString(i *int) string {
	if i != nil {
		return strconv.Itoa(*i)
	}
	return ""
}

// HTTPQueryEncodeInt will convert an Int pointer into a string and then add it to the query parameters of an HTTP
// request's URL.
func HTTPQueryEncodeInt(req *http.Request, key string, val *int) {
	if val != nil {
		httpQueryEncode(req, key, intPtrString(val))
	}
}

func int32PtrString(val *int32) string {
	if val != nil {
		return strconv.Itoa(int(*val))
	}
	return ""
}

func uint8PtrString(val *uint8) string {
	if val != nil {
		return strconv.Itoa(int(uint8(*val)))
	}
	return ""
}

// HTTPQueryEncodeInt32 will convert an Int32 pointer into a string and then add it to the query parameters of an HTTP
// request's URL.
func HTTPQueryEncodeInt32(req *http.Request, key string, val *int32) {
	if val != nil {
		httpQueryEncode(req, key, int32PtrString(val))
	}
}

// HTTPQueryEncodeString will convert an String pointer into a string and then add it to the query parameters of an HTTP
// request's URL.
func HTTPQueryEncodeString(req *http.Request, key string, val *string) {
	if val != nil {
		httpQueryEncode(req, key, *val)
	}
}

// HTTPQueryEncodeStringer will convert a stringer type into a string and then add it to the query parameters of
// an HTTP request's URL.
func HTTPQueryEncodeStringer(req *http.Request, key string, val stringer) {
	if val != nil {
		if str := val.String(); len(str) > 0 {
			httpQueryEncode(req, key, str)
		}
	}
}

// HTTPQueryEncodeStringSlice will convert a slice of strings into a string and then add it to the query parameters of
// an HTTP request's URL.
func HTTPQueryEncodeStrings(req *http.Request, key string, val []string) {
	if val != nil {
		slice := []string{}
		for _, v := range val {
			slice = append(slice, v)
		}
		httpQueryEncode(req, key, strings.Join(slice, ", "))
	}
}

// HTTPQueryEncodeTime will convert a time.Time type into a string and then add it to the query parameters of an HTTP
// request's URL.
func HTTPQueryEncodeTime(req *http.Request, key string, val *time.Time) {
	if val != nil {
		httpQueryEncode(req, key, val.String())
	}
}

// HTTPQueryEncodeUint8 will convert auint type into a string and then add it to the query parameters of an HTTP
// request's URL.
func HTTPQueryEncodeUint8(req *http.Request, key string, val *uint8) {
	if val != nil {
		httpQueryEncode(req, key, uint8PtrString(val))
	}
}

// parseErrorMessage takes a response and a status and builder an error message to send to the server.
func parseErrorMessage(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Errorf("Status Code %v: %v", resp.StatusCode, string(body))
}

// validateResponse is a switch condition that parses an error response
func validateResponse(res *http.Response) (err error) {
	if res == nil {
		err = fmt.Errorf("no response, check request and env file")
	} else {
		switch res.StatusCode {
		case
			http.StatusBadRequest,
			http.StatusUnauthorized,
			http.StatusInternalServerError,
			http.StatusNotFound,
			http.StatusForbidden:
			err = parseErrorMessage(res)
		}
	}
	return
}

func httpFetchRecursive(client http.Client, req *http.Request, opts Options, ep endpoint, params map[string]string,
	model interface{}, refetch bool) error {
	req.URL.Path = ep.Path(params)
	if opts != nil && !refetch {
		// if the request is to "refetch", then the options have already been encoded onto the request and we have no need
		// to re-encode.
		opts.EncodeQuery(req)
	}

	if req.Body != nil {
		req.Header.Set("content-type", "application/json")
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// If the request gets rate limited, then we need to re-fetch.
	if resp.StatusCode == http.StatusTooManyRequests {
		return httpFetchRecursive(client, req, opts, ep, params, model, true)
	}

	if err := validateResponse(resp); err != nil {
		return err
	}
	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)
	if model != nil {
		return json.NewDecoder(resp.Body).Decode(&model)
	}
	return nil
}

// HTTPFetch will make an HTTP request given a http.Client and a partially formatted http.Request, it will then try to
// edecode the model.
func HTTPFetch(client http.Client, req *http.Request, opts Options, ep endpoint, params map[string]string,
	model interface{}) error {
	return httpFetchRecursive(client, req, opts, ep, params, model, false)
}

// HTTPNewRequest will return a new request.  If the options are set, this function will encode a body if possible.
func HTTPNewRequest(method, url string, opts Options) (*http.Request, error) {
	if opts == nil {
		return http.NewRequest(method, url, nil)
	}
	buf, err := opts.EncodeBody()
	if err != nil {
		return nil, err
	}
	return http.NewRequest(method, url, buf)
}

// HTTPBodyFragment will add a new fragment to the HTTP request body.
func HTTPBodyFragment(body map[string]interface{}, key string, val interface{}) {
	if val != nil {
		body[key] = val
	}
}
