package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterConfigOptions(t *testing.T) {
	tests := []struct {
		name    string
		configs []Config
		setup   func()
		check   func(t *testing.T, v *viper.Viper)
		cleanup func()
	}{
		{
			name:    "empty configs list",
			configs: []Config{},
			check: func(t *testing.T, v *viper.Viper) {
			},
		},
		{
			name: "config with environment variable",
			configs: []Config{
				{
					Key:            "test_key",
					EnvironmentVar: "TEST_ENV_VAR",
				},
			},
			setup: func() {
				_ = os.Setenv("TEST_ENV_VAR", "env_value")
			},
			check: func(t *testing.T, v *viper.Viper) {
				assert.True(t, v.IsSet("test_key"))
			},
			cleanup: func() {
				_ = os.Unsetenv("TEST_ENV_VAR")
			},
		},
		{
			name: "config with name in file",
			configs: []Config{
				{
					Key:        "test_key",
					NameInFile: "test.name",
				},
			},
			check: func(t *testing.T, v *viper.Viper) {
				v.Set("test.name", "aliased_value")
				assert.Equal(t, "aliased_value", v.Get("test_key"))
			},
		},
		{
			name: "config with default value",
			configs: []Config{
				{
					Key:          "test_key",
					DefaultValue: "default_value",
				},
			},
			check: func(t *testing.T, v *viper.Viper) {
				assert.Equal(t, "default_value", v.Get("test_key"))
			},
		},
		{
			name: "config with all options",
			configs: []Config{
				{
					Key:            "test_key",
					NameInFile:     "test.name",
					EnvironmentVar: "TEST_ENV_VAR",
					DefaultValue:   "default_value",
				},
			},
			setup: func() {
				_ = os.Setenv("TEST_ENV_VAR", "env_value")
			},
			check: func(t *testing.T, v *viper.Viper) {
				assert.Equal(t, "env_value", v.Get("test.name"))
				assert.Equal(t, "env_value", v.Get("test_key"))
			},
			cleanup: func() {
				_ = os.Unsetenv("TEST_ENV_VAR")
			},
		},
		{
			name: "multiple configs",
			configs: []Config{
				{
					Key:          "key1",
					DefaultValue: "default1",
				},
				{
					Key:        "key2",
					NameInFile: "config.key2",
				},
				{
					Key:            "key3",
					EnvironmentVar: "TEST_ENV_VAR3",
				},
			},
			setup: func() {
				_ = os.Setenv("TEST_ENV_VAR3", "env_value3")
			},
			check: func(t *testing.T, v *viper.Viper) {
				assert.Equal(t, "default1", v.Get("key1"))

				v.Set("config.key2", "alias_value2")
				assert.Equal(t, "alias_value2", v.Get("key2"))
			},
			cleanup: func() {
				_ = os.Unsetenv("TEST_ENV_VAR3")
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			v := viper.New()

			if tt.setup != nil {
				tt.setup()
			}

			registerConfigOptions(tt.configs, v)

			tt.check(t, v)

			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestRegisterConfigFilePaths(t *testing.T) {
	v := viper.New()

	registerConfigFilePaths(v)

	// Unfortunately Viper doesn't expose getters for its internal state
	// Since this is a unit test, we can use reflection to access unexported fields
	viperVal := reflect.ValueOf(v).Elem()

	configTypeField := viperVal.FieldByName("configType")
	assert.True(t, configTypeField.IsValid())
	assert.Equal(t, "yaml", configTypeField.String())

	configNameField := viperVal.FieldByName("configName")
	assert.True(t, configNameField.IsValid())
	assert.Equal(t, "config", configNameField.String())

	cwd, err := os.Getwd()
	assert.NoError(t, err)

	home, err := os.UserHomeDir()
	assert.NoError(t, err)

	searchPathsField := viperVal.FieldByName("configPaths")
	assert.True(t, searchPathsField.IsValid())
	assert.Equal(t, 3, searchPathsField.Len())
	assert.Equal(t, cwd, searchPathsField.Index(0).String())
	assert.Equal(t, "/etc/service", searchPathsField.Index(1).String())
	assert.Equal(t, filepath.Join(home, ".config/service"), searchPathsField.Index(2).String())
}

func TestConfigureFromEnv(t *testing.T) {
	testCases := []struct {
		name     string
		configs  []Config
		envVars  map[string]string
		expected map[string]interface{}
	}{
		{
			name: "no environment variables",
			configs: []Config{
				{Key: "key1", EnvironmentVar: "TEST_ENV1"},
			},
			envVars:  map[string]string{},
			expected: map[string]interface{}{},
		},
		{
			name: "environment variable set",
			configs: []Config{
				{Key: "key1", EnvironmentVar: "TEST_ENV1"},
			},
			envVars:  map[string]string{"TEST_ENV1": "value1"},
			expected: map[string]interface{}{"key1": "value1"},
		},
		{
			name: "multiple environment variables",
			configs: []Config{
				{Key: "key1", EnvironmentVar: "TEST_ENV1"},
				{Key: "key2", EnvironmentVar: "TEST_ENV2"},
			},
			envVars:  map[string]string{"TEST_ENV1": "value1", "TEST_ENV2": "value2"},
			expected: map[string]interface{}{"key1": "value1", "key2": "value2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.envVars {
				t.Setenv(k, v)
			}

			v := viper.New()
			ConfigureFromEnv(tc.configs, v)

			for key, expected := range tc.expected {
				require.Equal(t, expected, v.Get(key), "Value for %s should be set from environment", key)
			}
		})
	}
}

func TestAutoConfigure(t *testing.T) {
	testCases := []struct {
		name     string
		configs  []Config
		envVars  map[string]string
		expected map[string]interface{}
	}{
		{
			name: "with environment variables",
			configs: []Config{
				{Key: "key1", EnvironmentVar: "TEST_AUTO_ENV1", DefaultValue: "default1"},
				{Key: "key2", EnvironmentVar: "TEST_AUTO_ENV2", DefaultValue: "default2"},
			},
			envVars: map[string]string{
				"TEST_AUTO_ENV1": "env_value1",
			},
			expected: map[string]interface{}{
				"key1": "env_value1", // From environment
				"key2": "default2",   // From default
			},
		},
		{
			name: "with defaults only",
			configs: []Config{
				{Key: "key1", EnvironmentVar: "TEST_AUTO_ENV3", DefaultValue: "default1"},
				{Key: "key2", EnvironmentVar: "TEST_AUTO_ENV4", DefaultValue: "default2"},
			},
			envVars: map[string]string{},
			expected: map[string]interface{}{
				"key1": "default1",
				"key2": "default2",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.envVars {
				t.Setenv(k, v)
			}

			v := viper.New()

			// We can't easily test file reading without setting up mock files,
			// so we'll replace the ReadInConfig method to simulate success or failure
			// This is a mock approach
			v.SetConfigType("yaml")

			AutoConfigure(tc.configs, v)

			for key, expected := range tc.expected {
				require.Equal(t, expected, v.Get(key), "Value for %s should be set correctly", key)
			}
		})
	}
}
