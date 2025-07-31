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
	ConfigPostgresHost          = "postgres_host"
	ConfigPostgresUser          = "postgres_user"
	ConfigPostgresPassword      = "postgres_password"
	ConfigPostgresDatabase      = "postgres_database"
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
		Key:          ConfigHTTPWriteTimeout,
		DefaultValue: 15 * time.Second,
	},
	{
		NameInFile:   "http.idle_timeout",
		Key:          ConfigHTTPIdleTimeout,
		DefaultValue: 60 * time.Second,
	},
	{
		NameInFile:   "postgres.host",
		Key:          ConfigPostgresHost,
		DefaultValue: "db",
	},
	{
		NameInFile:   "postgres.user",
		Key:          ConfigPostgresUser,
		DefaultValue: "postgres",
	},
	{
		NameInFile:   "postgres.password",
		Key:          ConfigPostgresPassword,
		DefaultValue: "postgres",
	},
	{
		NameInFile:   "postgres.database",
		Key:          ConfigPostgresDatabase,
		DefaultValue: "postgres",
	},
}
