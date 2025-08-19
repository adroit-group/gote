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
	ConfigValkeyHost            = "valkey_host"
	ConfigValkeyPort            = "valkey_port"
	ConfigValkeyDatabase        = "valkey_database"
	ConfigValkeyRetries         = "valkey_retries"
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
		NameInFile:     "valkey.host",
		EnvironmentVar: "VALKEY_HOST",
		Key:            ConfigValkeyHost,
		DefaultValue:   "valkey",
	},
	{
		NameInFile:     "valkey.port",
		EnvironmentVar: "VALKEY_PORT",
		Key:            ConfigValkeyPort,
		DefaultValue:   "6379",
	},
	{
		NameInFile:     "valkey.database",
		EnvironmentVar: "VALKEY_DATABASE",
		Key:            ConfigValkeyDatabase,
		DefaultValue:   0,
	},
	{
		NameInFile:     "valkey.retries",
		EnvironmentVar: "VALKEY_RETRIES",
		Key:            ConfigValkeyRetries,
		DefaultValue:   3,
	},
}
