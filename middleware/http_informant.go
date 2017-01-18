package middleware

import "time"

const (
	// HttpInformantKey is the key where an HttpInformer is stored in the Context's key/value store
	HttpInformantKey = "HttpInformant"
)

// HttpInformer is an interface for getting at information about an HTTP request,
// intended to be used by a logging middleware.
type HttpInformer interface {
	ID() string
	Message() string
	Method() string
	Path() string
	RemoteAddr() string

	SetID(string)
	SetMessage(string)
	SetMethod(string)
	SetPath(string)
	SetRemoteAddr(string)
}

// HttpInformant implements the HttpInformer interface
// for structured logging information. It is meant to be added to a request.Context
// or similar construct, for adding logging information at request and response level.
type HttpInformant struct {
	id         string
	remoteAddr string
	message    string
	method     string
	path       string
	duration   time.Duration
}

// ID is a unique value for the request, to match across request and response
func (hi *HttpInformant) ID() string {
	return hi.id
}

// RemoteAddr is the IP of the remote client
func (hi *HttpInformant) RemoteAddr() string {
	return hi.remoteAddr
}

// Method is the HTTP method (e.g. GET, POST, etc) associated with the request
func (hi *HttpInformant) Method() string {
	return hi.method
}

// Message is an arbitrary string to be added by the request handler
func (hi *HttpInformant) Message() string {
	return hi.message
}

// Path is the path requested
func (hi *HttpInformant) Path() string {
	return hi.path
}

// SetID sets the struct's id field
func (hi *HttpInformant) SetID(id string) {
	hi.id = id
}

// SetMessage sets the struct's message field
func (hi *HttpInformant) SetMessage(message string) {
	hi.message = message
}

// SetMethod sets the struct's method field
func (hi *HttpInformant) SetMethod(method string) {
	hi.method = method
}

// SetPath sets the struct's path method
func (hi *HttpInformant) SetPath(path string) {
	hi.path = path
}

// SetRemoteAddr sets the struct's remoteAddr field
func (hi *HttpInformant) SetRemoteAddr(remoteAddr string) {
	hi.remoteAddr = remoteAddr
}
