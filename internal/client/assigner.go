package client

import (
	"encoding/json"
	"io"
)

// Assigner is a struct that embeds a request, giving it access to all the request methods and
// data.  It serves as a safety mechanism, preventing a user from creating fetch chains in
// questionable order
type Assigner struct {
	*Request
	body io.ReadCloser
}

func newAssigner(req *Request) *Assigner {
	assigner := &Assigner{Request: req}
	return assigner
}

func (assigner *Assigner) decode(v interface{}) {
	if !assigner.errors.Any() {
		if err := json.NewDecoder(assigner.body).Decode(v); err != nil {
			assigner.errors.add(err)
		}
	}
}

func (assigner *Assigner) runAssignmentCallback(v interface{}) {
	if !assigner.errors.Any() && assigner.assignmentCallback != nil {
		if err := assigner.assignmentCallback(v, assigner.Request); err != nil {
			assigner.errors.add(err)
		}
	}
}

// Assign will assign the response body of an http request the interface v
func (assigner *Assigner) Assign(v interface{}) *Errors {
	defer assigner.body.Close()
	assigner.decode(v)
	assigner.runAssignmentCallback(v)
	return assigner.errors
}
