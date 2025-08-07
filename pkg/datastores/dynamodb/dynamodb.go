package dynamodb

import (
	"context"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Region   string `validate:"required"`
	Endpoint string `validate:"required"`
	Profile  string `validate:"required"`
	Retries  int    `validate:"required"`
}

func New(ctx context.Context, validate *validator.Validate, config Config) (*dynamodb.Client, error) {
	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	cfg, err := awsconfig.LoadDefaultConfig(
		ctx,
		awsconfig.WithRegion(config.Region),
		awsconfig.WithBaseEndpoint(config.Endpoint),
		awsconfig.WithSharedConfigProfile(config.Profile),
		awsconfig.WithRetryMaxAttempts(config.Retries),
	)
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}
