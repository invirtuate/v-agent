// Package config ensures the app is configured properly
package config

import "errors"

var (
	// ErrConfigNotInitialized returned if the configuration is not initialized
	ErrConfigNotInitialized = errors.New("config not initialized")
	// ErrLabelNotExist returned when the specified label doesnt exist
	ErrLabelNotExist = errors.New("label does not exist")

	ErrPortInvalid     = errors.New("port is invalid")
	ErrIntervalInvalid = errors.New("interval is invalid")
	ErrMissingScheme   = errors.New("missing http/https")

	ErrNotInK8s = errors.New("not running in kubernetes")

	ErrKubernetesNamespaceInvalid  = errors.New("namespace(s) invalid")
	ErrKubernetesNamespaceNotSet   = errors.New("namespace not set")
	ErrKubernetesNamespaceNotExist = errors.New("namespace does not exist")

	ErrSMARTDeviceNotExist = errors.New("smart.block_device does not exist")

	ErrDCGMEndpointNotSet   = errors.New("dcgm.endpoint not set")
	ErrDCGMEndpointNotExist = errors.New("dcgm.endpoint does not exist")
)
