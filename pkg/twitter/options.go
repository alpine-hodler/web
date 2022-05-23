package twitter

import (
	"io"
	"net/http"

	"github.com/alpine-hodler/sdk/pkg/scalar"
	"github.com/alpine-hodler/sdk/tools"
)

// * This is a generated file, do not edit

// ComplianceJobsOptions are options for API requests.
type ComplianceJobsOptions struct {
	// Status allows to filter by job status. Only one filter can be specified per request. Default: `all`
	Status *scalar.Status `json:"status" bson:"status"`

	// Type allows to filter by job type - either by tweets or user ID. Only one filter (tweets or users) can be specified
	// per request.
	Type scalar.ComplianceJob `json:"type" bson:"type"`
}

// AllTweetsOptions are options for API requests.
type AllTweetsOptions struct {
	// MaxResults are the maximum number of search results to be returned by a request. A number between 10 and the system
	// limit (currently 500). By default, a request response will return 10 results.
	MaxResults *int `json:"max_results" bson:"max_results"`
}

// TweetsOptions are options for API requests.
type TweetsOptions struct {
	// TODO
	Ids []string `json:"ids" bson:"ids"`
}

func (opts *AllTweetsOptions) EncodeBody() (buf io.Reader, err error)      { return }
func (opts *ComplianceJobsOptions) EncodeBody() (buf io.Reader, err error) { return }
func (opts *TweetsOptions) EncodeBody() (buf io.Reader, err error)         { return }

// SetMaxResults sets the MaxResults field on AllTweetsOptions. MaxResults are the maximum number of search results to
// be returned by a request. A number between 10 and the system limit (currently 500). By default, a request response
// will return 10 results.
func (opts *AllTweetsOptions) SetMaxResults(MaxResults int) *AllTweetsOptions {
	opts.MaxResults = &MaxResults
	return opts
}

// SetType sets the Type field on ComplianceJobsOptions. Type allows to filter by job type - either by tweets or user
// ID. Only one filter (tweets or users) can be specified per request.
func (opts *ComplianceJobsOptions) SetType(Type scalar.ComplianceJob) *ComplianceJobsOptions {
	opts.Type = Type
	return opts
}

// SetStatus sets the Status field on ComplianceJobsOptions. Status allows to filter by job status. Only one filter can
// be specified per request. Default: `all`
func (opts *ComplianceJobsOptions) SetStatus(Status scalar.Status) *ComplianceJobsOptions {
	opts.Status = &Status
	return opts
}

// SetIds sets the Ids field on TweetsOptions. TODO
func (opts *TweetsOptions) SetIds(Ids []string) *TweetsOptions {
	opts.Ids = Ids
	return opts
}

func (opts *AllTweetsOptions) EncodeQuery(req *http.Request) {
	if opts != nil {
		tools.HTTPQueryEncodeInt(req, "max_results", opts.MaxResults)
	}
	return
}

func (opts *ComplianceJobsOptions) EncodeQuery(req *http.Request) {
	if opts != nil {
		tools.HTTPQueryEncodeStringer(req, "type", &opts.Type)
		tools.HTTPQueryEncodeStringer(req, "status", opts.Status)
	}
	return
}

func (opts *TweetsOptions) EncodeQuery(req *http.Request) {
	if opts != nil {
		tools.HTTPQueryEncodeStrings(req, "ids", opts.Ids)
	}
	return
}
