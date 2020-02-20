package model

// Method represents the int of the http method
type Method int

// GET, POST and PUT http methods
const (
	GET Method = iota + 1
	POST
	PUT
)

// MethodFromString returns an http method from the method string
func MethodFromString(method string) Method {
	switch method {
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	}

	return -1
}

// Request represents an HTTP request
type Request struct {
	Path       string
	UUID       string
	Method     Method
	RemoteAddr string
}

// NewRequest returns a new request
func NewRequest(path string, UUID string, method Method, remoteAddr string) Request {
	return Request{
		Path:       path,
		UUID:       UUID,
		Method:     method,
		RemoteAddr: remoteAddr,
	}
}
