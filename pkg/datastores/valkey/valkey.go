package valkey

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/valkey-io/valkey-go"
	"log/slog"
	"net"
	"time"
)

var (
	ErrFailedValkeyConnection = errors.New("failed valkey connection")
)

// Config represents a configuration option.
type Config struct {
	// Host is the host of the Valkey server.
	Host string `validate:"required"`
	// Port is the port of the Valkey server.
	Port string `validate:"required,numeric"`
	// Database is the number of the Valkey database.
	Database int
	// Retries defines how many times should we try to connect.
	Retries int `validate:"required"`
}

// New validates the configuration and tries to connect to the Valkey database.
func New(validate *validator.Validate, config Config) (valkey.Client, error) {
	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	option := valkey.ClientOption{
		InitAddress: []string{net.JoinHostPort(config.Host, config.Port)},
		SelectDB:    config.Database,
		Dialer:      net.Dialer{Timeout: 15 * time.Second},
	}

	for i := 0; i < config.Retries; i++ {
		conn, err := valkey.NewClient(option)
		if err != nil {
			slog.Error(ErrFailedValkeyConnection.Error(), "error", err)
			continue
		}

		slog.Info("successful valkey connection", "connection string", fmt.Sprintf("%s:%s/%d", config.Host, config.Port, config.Database))
		return conn, nil
	}

	return nil, ErrFailedValkeyConnection
}
