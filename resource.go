package resourcesql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
)

type Resource struct {
	*sql.DB

	name               string
	config             Config
	skipTestWhenStarts bool
}

type Option = func(service *Resource)

const (
	defaultName = "SQL"
)

// NewResource returns a new instance of the SQLResourceService interface.
func NewResource(options ...Option) *Resource {
	r := &Resource{
		name: defaultName,
	}
	for _, opt := range options {
		opt(r)
	}
	return r
}

// WithName is an option that will set the service name for a SQLResourceService.
func WithName(name string) Option {
	return func(service *Resource) {
		service.name = name
	}
}

// WithConfig is an option that will set the config for a SQLResourceService.
func WithConfig(config Config) Option {
	return func(service *Resource) {
		service.config = config
	}
}

// WithSkipTestWhenStarts flags the SQLResourceService to not test the connection at the initialization.
func WithSkipTestWhenStarts(value bool) Option {
	return func(service *Resource) {
		service.skipTestWhenStarts = value
	}
}

// Name will return a human identifiable name for this service. Ex: Postgresql Connection.
func (service *Resource) Name() string {
	return service.name
}

// Start will start the service in a blocking way.
//
// If the service is successfully started, `nil` should be returned. Otherwise, an error must be returned.
func (service *Resource) Start(_ context.Context) error {
	if service.config == nil {
		return ErrMissingConfiguration
	}
	dsn := service.config.GetDSN()

	u, err := url.Parse(dsn)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidDSN, err.Error())
	}

	db, err := sql.Open(u.Scheme, dsn)
	if err != nil {
		return err
	}

	if maxIdleConnections := service.config.GetMaxIdleConnections(); maxIdleConnections > 0 {
		db.SetMaxIdleConns(maxIdleConnections)
	}

	if maxOpenConns := service.config.GetMaxOpenConns(); maxOpenConns > 0 {
		db.SetMaxOpenConns(maxOpenConns)
	}

	if connMaxLifetime := service.config.GetConnMaxLifetime(); connMaxLifetime > 0 {
		db.SetConnMaxLifetime(connMaxLifetime)
	}

	if !service.skipTestWhenStarts {
		err = db.Ping()
		if err != nil {
			_ = db.Close()
			return fmt.Errorf("%w: %s", ErrInitialConnectionTestFailed, err.Error())
		}
	}

	service.DB = db
	return nil
}

// Stop will stop this service.
//
// For most implementations it will be blocking and should return only when the service finishes stopping.
//
// If the service is successfully stopped, `nil` should be returned. Otherwise, an error must be returned.
func (service *Resource) Stop(_ context.Context) error {
	return service.DB.Close()
}
