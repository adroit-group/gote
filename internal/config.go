package internal

import (
	"time"

	"github.com/adroit-group/gote/pkg/config"
)

const (
	ConfigHTTPBasePath          = "base_path"
	ConfigHTTPPort              = "http_port"
	ConfigHTTPReadTimeout       = "http_read_timeout"
	ConfigHTTPReadHeaderTimeout = "http_read_header_timeout"
	ConfigHTTPWriteTimeout      = "http_write_timeout"
	ConfigHTTPIdleTimeout       = "http_idle_timeout"
	ConfigDynamoDBRegion        = "dynamo_db_region"
	ConfigDynamoDBEndpoint      = "dynamo_db_endpoint"
	ConfigDynamoDBProfile       = "dynamo_db_profile"
	ConfigDynamoDBRetries       = "dynamo_db_retries"
)

var Configuration = []config.Config{
	{
		NameInFile:     "http.base_path",
		EnvironmentVar: "HTTP_BASE_PATH",
		Key:            ConfigHTTPBasePath,
		DefaultValue:   "/api",
	},
	{
		NameInFile:   "http.port",
		Key:          ConfigHTTPPort,
		DefaultValue: 80,
	},
	{
		NameInFile:   "http.read_timeout",
		Key:          ConfigHTTPReadTimeout,
		DefaultValue: 15 * time.Second,
	},
	{
		NameInFile:   "http.read_header_timeout",
		Key:          ConfigHTTPReadHeaderTimeout,
		DefaultValue: 15 * time.Second,
	},
	{
		NameInFile:   "http.write_timeout",
		Key:          "ConfigHTTPWriteTimeout",
		DefaultValue: 15 * time.Second,
	},
	{
		NameInFile:   "http.idle_timeout",
		Key:          ConfigHTTPIdleTimeout,
		DefaultValue: 60 * time.Second,
	},
	{
		NameInFile:     "dynamodb.region",
		EnvironmentVar: "DYNAMO_DB_REGION",
		Key:            ConfigDynamoDBRegion,
		DefaultValue:   "eu-west-1",
	},
	{
		NameInFile:     "dynamodb.endpoint",
		EnvironmentVar: "DYNAMO_DB_ENDPOINT",
		Key:            ConfigDynamoDBEndpoint,
		DefaultValue:   "http://dynamodb:8000",
	},
	{
		NameInFile:     "dynamodb.profile",
		EnvironmentVar: "DYNAMO_DB_PROFILE",
		Key:            ConfigDynamoDBProfile,
		DefaultValue:   "default",
	},
	{
		NameInFile:     "dynamodb.retries",
		EnvironmentVar: "DYNAMO_DB_RETRIES",
		Key:            ConfigDynamoDBRetries,
		DefaultValue:   3,
	},
}
