package client

// Common HTTP methods.
//
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
type Method uint8

const (
	GET Method = iota
	POST
	PUT
	DELETE
)

// String will return the string expectation of the method
func (m Method) String() string {
	return map[Method]string{
		GET:    "GET",
		POST:   "POST",
		PUT:    "PUT",
		DELETE: "DELETE",
	}[m]
}
