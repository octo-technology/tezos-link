package errors

import "errors"

var ErrProjectNotFound = errors.New("project not found")
var ErrNoProxyResponse = errors.New("no response from proxy")
