package model

// Health represents to connection to a database
type Health struct {
	ConnectedToDB bool `json:"connectedToDb"`
}

// NewHealth returns a new health object
func NewHealth(connectedToDB bool) Health {
	return Health{
		ConnectedToDB: connectedToDB,
	}
}
