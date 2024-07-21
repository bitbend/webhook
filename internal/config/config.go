package config

import (
	"bytes"
	_ "embed"
	"errors"
	"github.com/spf13/viper"
	"log"
	"strings"
)

const ConfigurationPrefix = "DRIFTHOOK"

type ServerConfiguration struct {
	Host string `mapstructure:"host"`
	Port uint   `mapstructure:"port"`
	SSL  bool   `mapstructure:"ssl"`
	Env  string `mapstructure:"env"`
}

type PostgresUserConfiguration struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type PostgresConfiguration struct {
	Host              string                    `mapstructure:"host"`
	Port              uint                      `mapstructure:"port"`
	User              PostgresUserConfiguration `mapstructure:"user"`
	Database          string                    `mapstructure:"database"`
	Options           string                    `mapstructure:"options"`
	MaxOpenConnection uint                      `mapstructure:"max_open_connection"`
	MaxIdleConnection uint                      `mapstructure:"max_idle_connection"`
}

type DatabaseConfiguration struct {
	Postgres PostgresConfiguration `mapstructure:"postgres"`
}

type NATSConfiguration struct {
	Urls     []string `mapstructure:"urls"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
}

type QueueConfiguration struct {
	NATS NATSConfiguration `mapstructure:"nats"`
}

type S3Configuration struct {
	AccessKey      string `mapstructure:"access_key"`
	SecretKey      string `mapstructure:"secret_key"`
	Bucket         string `mapstructure:"bucket"`
	PathPrefix     string `mapstructure:"path_prefix"`
	Region         string `mapstructure:"region"`
	ForcePathStyle bool   `mapstructure:"force_path_style"`
	DisableSSL     bool   `mapstructure:"disable_ssl"`
}

type StorageConfiguration struct {
	S3 S3Configuration `mapstructure:"s3"`
}

type Configuration struct {
	Server   ServerConfiguration   `mapstructure:"server"`
	Database DatabaseConfiguration `mapstructure:"database"`
	Queue    QueueConfiguration    `mapstructure:"queue"`
	Storage  StorageConfiguration  `mapstructure:"storage"`
}

//go:embed default.yaml
var defaultConfiguration []byte

func LoadConfig(path string, name string) (*Configuration, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName(name)
	viper.AddConfigPath(path)

	viper.AutomaticEnv()
	viper.SetEnvPrefix(ConfigurationPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		log.Fatalf("failed to load default config: %v", err)
	}

	if err := viper.MergeInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Printf("%s.yaml config file not found at %s", name, path)
		} else {
			log.Fatalf("failed to merge config: %v", err)
		}
	}

	var configuration Configuration
	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("failed to unmarshall configuration: %v", err)
	}

	return &configuration, nil
}
