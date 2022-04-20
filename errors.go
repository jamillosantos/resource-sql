package resourcesql

import "github.com/pkg/errors"

var (
	// ErrMissingConfiguration is returned when no configuration is given to the SQLResourceService implementation.
	ErrMissingConfiguration = errors.New("config not provided")

	// ErrInitialConnectionTestFailed is returned when the start up connection test fails.
	ErrInitialConnectionTestFailed = errors.New("initial connection test failed")

	// ErrInvalidDSN is returned when the start up connection test fails.
	ErrInvalidDSN = errors.New("invalid dsn")
)
