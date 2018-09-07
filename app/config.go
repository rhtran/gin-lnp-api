package app

import (
	"fmt"

	"os"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
	"time"
)

// Config stores the application-wide configurations
var Config appConfig

type appConfig struct {
	// the path to the error message file. Defaults to "config/errors.yaml"
	ErrorFile string `mapstructure:"error_file"`
	Hostname  string `mapstructure:"hostname"`
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// the data source name (DSN) for connecting to the database. required.
	CassHost        string `mapstructure:"cassandra_host"`
	ClusterUsername string `mapstructure:"cassandra_username"`
	ClusterPassword string `mapstructure:"cassandra_password"`
	Keyspace        string `mapstructure:"cassandra_keyspace"`
	DataCenter		string `mapstructure:"cassandra_data_center"`
	Timeout			time.Duration 	`mapstructure:"cassandra_timeout"`
	ProtoVersion	int 	`mapstructure:"cassandra_proto_version"`

}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.CassHost, validation.Required),
		validation.Field(&config.ClusterUsername, validation.Required),
		validation.Field(&config.ClusterPassword, validation.Required),
		validation.Field(&config.Keyspace, validation.Required),
		validation.Field(&config.DataCenter, validation.Required),
		validation.Field(&config.Timeout, validation.Required),
		validation.Field(&config.ProtoVersion, validation.Required),
	)
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "LNP_" in their names are also read automatically.
func LoadConfig(configPaths ...string) error {
	v := viper.New()

	if os.Getenv("ENV") == "PROD" {
		v.SetConfigName("app-prod")
	} else if os.Getenv("ENV") == "STG" {
		v.SetConfigName("app-stg")
	} else if os.Getenv("ENV") == "DEV" {
		v.SetConfigName("app-dev")
	} else {
		v.SetConfigName("app")
	}

	v.SetConfigType("yaml")
	v.SetEnvPrefix("lnp")
	v.AutomaticEnv()
	v.SetDefault("error_file", "config/errors.yaml")
	v.SetDefault("server_port", 8080)
	v.SetDefault("jwt_signing_method", "HS256")

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return Config.Validate()
}
