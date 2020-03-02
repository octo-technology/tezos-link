package model

// Action represents the int
type Action int

// OBTAIN, PUSH and MODIFY methods
const (
	OBTAIN Action = iota + 1
	PUSH
	MODIFY
)

// Request represents an request made to the proxy
type Request struct {
	Path       string
	UUID       string
	Action     Action
	RemoteAddr string
}

// NewRequest returns a new request
func NewRequest(path string, UUID string, action Action, remoteAddr string) Request {
	return Request{
		Path:       path,
		UUID:       UUID,
		Action:     action,
		RemoteAddr: remoteAddr,
	}
}
