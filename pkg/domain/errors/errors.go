package errors

import "errors"

var ErrProjectNotFound = errors.New("project not found")
var ErrNoProjectName = errors.New("project name not defined")
var ErrNoProxyResponse = errors.New("no response from proxy")
