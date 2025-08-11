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

type Config struct {
	Host     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
	Retries  int    `validate:"required"`
}

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
