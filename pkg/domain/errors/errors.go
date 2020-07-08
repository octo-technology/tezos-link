package errors

import "errors"

// ErrProjectNotFound when a project is not found
var ErrProjectNotFound = errors.New("project not found")

// ErrNoProjectName when a project has no name
var ErrNoProjectName = errors.New("project name not defined")

// ErrNoProxyResponse when the proxy doesn't respond
var ErrNoProxyResponse = errors.New("no response from proxy")

var ErrNoMetricsFound = errors.New("no old metrics")
