package config

import (
	"bytes"
	_ "embed"
	"errors"
	"github.com/spf13/viper"
	"log"
	"strings"
)

const EnvPrefix = "DRIFTBASE_WEBHOOK"

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port uint   `mapstructure:"port"`
	SSL  bool   `mapstructure:"ssl"`
	Env  string `mapstructure:"env"`
}

type UserConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type PostgresConfig struct {
	Host               string     `mapstructure:"host"`
	Port               uint       `mapstructure:"port"`
	User               UserConfig `mapstructure:"user"`
	Database           string     `mapstructure:"database"`
	Options            string     `mapstructure:"options"`
	MaxOpenConnections uint       `mapstructure:"max_open_connections"`
	MaxIdleConnections uint       `mapstructure:"max_idle_connections"`
}

type YugabyteConfig struct {
	Host               string     `mapstructure:"host"`
	Port               uint       `mapstructure:"port"`
	User               UserConfig `mapstructure:"user"`
	Database           string     `mapstructure:"database"`
	Options            string     `mapstructure:"options"`
	MaxOpenConnections uint       `mapstructure:"max_open_connections"`
	MaxIdleConnections uint       `mapstructure:"max_idle_connections"`
}

type RedisConfig struct {
	Url      string `mapstructure:"url"`
	Database string `mapstructure:"database"`
	Password string `mapstructure:"password"`
}

type NATSConfig struct {
	Urls     []string `mapstructure:"urls"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
}

type S3Config struct {
	Endpoint       string `mapstructure:"endpoint"`
	AccessKey      string `mapstructure:"access_key"`
	SecretKey      string `mapstructure:"secret_key"`
	Bucket         string `mapstructure:"bucket"`
	PathPrefix     string `mapstructure:"path_prefix"`
	Region         string `mapstructure:"region"`
	ForcePathStyle bool   `mapstructure:"force_path_style"`
	DisableSSL     bool   `mapstructure:"disable_ssl"`
}

type DatabaseConfig struct {
	Provider string         `mapstructure:"provider"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Yugabyte YugabyteConfig `mapstructure:"yugabyte"`
}

type CacheConfig struct {
	Redis RedisConfig `mapstructure:"redis"`
}

type TaskQueueConfig struct {
	NATS NATSConfig `mapstructure:"nats"`
}

type StorageConfig struct {
	S3 S3Config `mapstructure:"s3"`
}

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Database  DatabaseConfig  `mapstructure:"database"`
	Cache     CacheConfig     `mapstructure:"cache"`
	TaskQueue TaskQueueConfig `mapstructure:"task_queue"`
	Storage   StorageConfig   `mapstructure:"storage"`
}

//go:embed default.yaml
var defaultConfig []byte

func NewConfig(path string, name string) *Config {
	return loadConfig(path, name)
}

func loadConfig(path string, name string) *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigName(name)
	viper.AddConfigPath(path)

	viper.AutomaticEnv()
	viper.SetEnvPrefix(EnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
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

	var configuration Config
	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("failed to unmarshall configuration: %v", err)
	}

	return &configuration
}
