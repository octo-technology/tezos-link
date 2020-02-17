package model

type Health struct {
    ConnectedToDB   bool `json:"connectedToDb"`
}

func NewHealth(connectedToDB bool) Health {
    return Health{
        ConnectedToDB: connectedToDB,
    }
}
