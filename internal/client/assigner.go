package client

import (
	"encoding/json"
	"io"
	"time"
)

// Assigner is a struct that embeds a request, giving it access to all the request methods and data.  It serves as a
// safety mechanism, preventing a user from creating fetch chains in questionable order
type Assigner struct {
	*Request
	body io.ReadCloser
}

func newAssigner(req *Request) *Assigner {
	assigner := &Assigner{Request: req}
	return assigner
}

// Profile represents a profile to interact with the API.
type Profile struct {
	Active    bool      `json:"active" bson:"active"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	HasMargin bool      `json:"has_margin" bson:"has_margin"`
	Id        string    `json:"id" bson:"id"`
	IsDefault bool      `json:"is_default" bson:"is_default"`
	Name      string    `json:"name" bson:"name"`
	UserId    string    `json:"user_id" bson:"user_id"`
}

func (assigner *Assigner) decode(v interface{}) {
	if !assigner.errors.Any() {
		// TODO: wrap these debugging tools into the assigner as a logging method, or do something like that.
		// body, err := ioutil.ReadAll(assigner.body)
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("Raw Response Body:\n%v\n", string(body))

		// var data map[string]interface{}
		// if err := json.Unmarshal(body, &data); err != nil {
		// 	panic(err)
		// }
		// fmt.Println("DATA:", data)

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
	defer func() {
		if body := assigner.body; body != nil {
			body.Close()
		}
	}()
	assigner.decode(v)
	assigner.runAssignmentCallback(v)
	return assigner.errors
}

func (assigner *Assigner) NoAssignment() *Errors {
	defer func() {
		if body := assigner.body; body != nil {
			body.Close()
		}
	}()
	assigner.runAssignmentCallback(nil)
	return assigner.errors
}
