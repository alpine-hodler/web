package tools

import (
	"net/url"
	"strings"
)

// URIBuilderComponents are the classification for which components of a URI we are trying to store in memory to
// prepare for an HTTP request.
type URIBuilderComponent uint8

const (
	// A path. The path identifies the specific resource in the host that the web client wants to access. For example,
	// /software/htp/cics/index.html.
	URIBuilderComponentPath URIBuilderComponent = iota

	// If a query string is used, it follows the path component, and provides a string of information that the resource
	// can use for some purpose (for example, as parameters for a search or as data to be processed). The query string is
	// usually a string of name and value pairs; for example, term=bluebird. Name and value pairs are separated from each
	// other by an ampersand (&); for example, term=bluebird&source=browser-search.
	URIBuilderComponentQuery
)

type uriBuilderComponentMap map[string]string

// URIBuilder maps componenets to their key/value input in a URI.
type URIBuilder map[URIBuilderComponent]uriBuilderComponentMap

// URIBuilderAccessor is an interface to build sections of a URI.
type URIBuilderAccessor interface {
	// PostAuthority is everything after the authority portion of a URI.  This includes path, filename, param matrix,
	// query and fragment.  See https://www.rfc-editor.org/rfc/rfc3986#section-3 for more details.
	PostAuthority(URIBuilder) string
}

// Add will create the tree for a component and a key/value pair in memory.
func (uriBuilder URIBuilder) Add(component URIBuilderComponent, key, value string) {
	if uriBuilder[component] == nil {
		uriBuilder[component] = make(uriBuilderComponentMap)
	}
	uriBuilder[component][key] = value
}

// Get will return the value for a component+key.
func (uriBuilder URIBuilder) Get(component URIBuilderComponent, key string) (value string) {
	if componentMap := uriBuilder[component]; componentMap != nil {
		value = componentMap[key]
	}
	return
}

// QueryPath get the query path from a URIBuilder.
func (uriBuilder URIBuilder) QueryPath() *url.URL {
	u, _ := url.Parse("")
	q, _ := url.ParseQuery(u.RawQuery)
	for key, queryParam := range uriBuilder[URIBuilderComponentQuery] {
		q.Add(key, queryParam)
	}
	u.RawQuery = q.Encode()
	return u
}

func JoinEndpointParts(parts ...string) string {
	return "/" + strings.Join(parts, "/")
}
