package datastores

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/url"
)

type PostgresConfig struct {
	Host     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
}

func NewPostgres(ctx context.Context, validate *validator.Validate, config PostgresConfig) (*pgxpool.Pool, error) {
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

	return pgxpool.New(ctx, pgURL.String())
}
