package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/url"
)

var (
	ErrFailedPostgresConnection = errors.New("failed postgres connection")
)

// Config represents a configuration option.
type Config struct {
	// Host is the host of the Postgresql server.
	Host string `validate:"required"`
	// User is the Postgresql user.
	User string `validate:"required"`
	// Password is the password for the Postgresql user.
	Password string `validate:"required"`
	// Database is the Postgresql database name.
	Database string `validate:"required"`
	// Retries defines how many times should we try to connect.
	Retries int `validate:"required"`
}

// New validates the configuration and tries to connect to the Postgresql database.
func New(ctx context.Context, validate *validator.Validate, config Config) (*pgxpool.Pool, error) {
	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	pgURL := url.URL{}
	pgURL.Scheme = "postgres"
	pgURL.User = url.UserPassword(config.User, config.Password)
	pgURL.Host = config.Host
	pgURL.Path = config.Database

	values := url.Values{}
	values.Set("sslmode", "disable")
	pgURL.RawQuery = values.Encode()

	for i := 0; i < config.Retries; i++ {
		conn, err := pgxpool.New(ctx, pgURL.String())
		if err != nil {
			slog.Error(ErrFailedPostgresConnection.Error(), "error", err)
			continue
		}
		if err = conn.Ping(ctx); err != nil {
			slog.Error(ErrFailedPostgresConnection.Error(), "error", err)
			continue
		}

		slog.Info("successful postgres connection", "connection string", fmt.Sprintf("%s:5432/%s", config.Host, config.Database))
		return conn, nil
	}

	return nil, ErrFailedPostgresConnection
}
