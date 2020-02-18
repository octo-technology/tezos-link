package model

type Method int

const(
    GET Method = iota + 1
    POST
    PUT
)

func MethodFromString(method string) Method {
    switch method {
        case "GET":     return GET
        case "POST":    return POST
        case "PUT":     return PUT
    }

    return -1
}

type Request struct {
    Path        string
    UUID        string
    Method      Method
    RemoteAddr  string
}

func NewRequest(path string, UUID string, method Method, remoteAddr string) Request {
    return Request{
        Path:   path,
        UUID:   UUID,
        Method: method,
        RemoteAddr:  remoteAddr,
    }
}
