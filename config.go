package resourcesql

import (
	"time"
)

// Config represents the configuration for a SQLResourceService.
type Config interface {
	GetDSN() string
	GetMaxIdleConnections() int
	GetMaxOpenConns() int
	GetConnMaxLifetime() time.Duration
}

type PlatformConfig struct {
	DSN                string        `config:"dsn,secret,required"`
	MaxIdleConnections int           `config:"max_idle_connections"`
	MaxOpenConns       int           `config:"max_open_conns"`
	ConnMaxLifetime    time.Duration `config:"conn_max_lifetime"`
}

func (p *PlatformConfig) GetDSN() string {
	return p.DSN
}

func (p *PlatformConfig) GetMaxIdleConnections() int {
	return p.MaxIdleConnections
}

func (p *PlatformConfig) GetMaxOpenConns() int {
	return p.MaxOpenConns
}

func (p *PlatformConfig) GetConnMaxLifetime() time.Duration {
	return p.ConnMaxLifetime
}
