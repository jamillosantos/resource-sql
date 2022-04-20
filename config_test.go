package resourcesql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPlatformConfig(t *testing.T) {
	wantDSN := "dsn"
	wantMaxIdleConnections := 1
	wantMaxOpenConns := 2
	wantConnMaxLifetime := time.Duration(3)

	config := PlatformConfig{
		DSN:                wantDSN,
		MaxIdleConnections: wantMaxIdleConnections,
		MaxOpenConns:       wantMaxOpenConns,
		ConnMaxLifetime:    wantConnMaxLifetime,
	}
	assert.Equal(t, wantDSN, config.GetDSN())
	assert.Equal(t, wantMaxIdleConnections, config.GetMaxIdleConnections())
	assert.Equal(t, wantMaxOpenConns, config.GetMaxOpenConns())
	assert.Equal(t, wantConnMaxLifetime, config.GetConnMaxLifetime())
}
