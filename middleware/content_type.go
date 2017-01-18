// ContentType ensures that the "Content-Type" HTTP header is always "application/json"
// in all responses from the server.
package middleware

import (
	"net/http"

	"github.com/urfave/negroni"
)

const (
	// The HTTP header that will be altered
	HttpContentTypeKey = "Content-Type"
	// The value that will go in the header
	HttpContentTypeValue = "application/json"
)

// ContentType Struct has no members b/c it's just a way to satisfy the interface that Negroni requires for middleware
type ContentType struct {
}

// NewContentType returns a new instance of ContentType
func NewContentType() *ContentType {
	return &ContentType{}
}

func (r *ContentType) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	res := w.(negroni.ResponseWriter)
	res.Before(func(res negroni.ResponseWriter) {
		w.Header().Set(HttpContentTypeKey, HttpContentTypeValue)
	})
	next(w, req)
}
