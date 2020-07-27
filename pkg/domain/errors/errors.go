package errors

import "errors"

// ErrProjectNotFound when a project is not found
var ErrProjectNotFound = errors.New("project not found")

// ErrNoProjectName when a project has no name
var ErrNoProjectName = errors.New("project name not defined")

// ErrNoProxyResponse when the proxy doesn't respond
var ErrNoProxyResponse = errors.New("no response from proxy")

// ErrNoMetricsFound when there is no metrics found
var ErrNoMetricsFound = errors.New("no old metrics")

// ErrInvalidNetwork when there is an invalid network
var ErrInvalidNetwork = errors.New("invalid network")
