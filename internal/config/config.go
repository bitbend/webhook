package config

type ServerConfiguration struct {
	Host string `json:"host" yaml:"host" mapstructure:"DRIFTHOOK_SERVER_HOST"`
	Port uint   `json:"port" yaml:"port" mapstructure:"DRIFTHOOK_SERVER_PORT"`
	SSL  bool   `json:"ssl" yaml:"ssl" mapstructure:"DRIFTHOOK_SERVER_SSL"`
	Env  string `json:"env" yaml:"env" mapstructure:"DRIFTHOOK_SERVER_ENV"`
}

type PostgresUserConfiguration struct {
	Username string `json:"username" yaml:"username" mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_USER_USERNAME"`
	Password string `json:"password" yaml:"password" mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_USER_PASSWORD"`
}

type PostgresConfiguration struct {
	Host              string                    `json:"host" yaml:"host" mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_HOST"`
	Port              uint                      `json:"port" yaml:"port" mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_PORT"`
	User              PostgresUserConfiguration `json:"user" yaml:"user" mapstructure:",squash"`
	DatabaseName      string                    `json:"database_name" yaml:"databaseName" mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_DATABASE_NAME"`
	Options           string                    `json:"options" yaml:"options" mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_OPTIONS"`
	MaxOpenConnection uint                      `json:"max_open_connection" yaml:"maxOpenConnection" mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_MAX_OPEN_CONNECTION"`
	MaxIdleConnection uint                      `json:"max_idle_connection" yaml:"maxIdleConnection" mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_MAX_IDLE_CONNECTION"`
}

type DatabaseConfiguration struct {
	Postgres PostgresConfiguration `json:"postgres" yaml:"postgres" mapstructure:",squash"`
}

type NATSConfiguration struct {
	Urls     []string `json:"urls" yaml:"urls" mapstructure:"DRIFTHOOK_QUEUE_NATS_URLS"`
	Username string   `json:"username" yaml:"username" mapstructure:"DRIFTHOOK_QUEUE_NATS_USERNAME"`
	Password string   `json:"password" yaml:"password" mapstructure:"DRIFTHOOK_QUEUE_NATS_PASSWORD"`
}

type QueueConfiguration struct {
	NATS NATSConfiguration `json:"nats" yaml:"nats" mapstructure:",squash"`
}

type S3Configuration struct {
	AccessKey      string `json:"access_key" yaml:"accessKey" mapstructure:"DRIFTHOOK_STORAGE_S3_ACCESS_KEY"`
	SecretKey      string `json:"secret_key" yaml:"secretKey" mapstructure:"DRIFTHOOK_STORAGE_S3_SECRET_KEY"`
	BucketName     string `json:"bucket_name" yaml:"bucketName" mapstructure:"DRIFTHOOK_STORAGE_S3_BUCKET_NAME"`
	PathPrefix     string `json:"path_prefix" yaml:"pathPrefix" mapstructure:"DRIFTHOOK_STORAGE_S3_PATH_PREFIX"`
	Region         string `json:"region" yaml:"region" mapstructure:"DRIFTHOOK_STORAGE_S3_REGION"`
	ForcePathStyle bool   `json:"force_path_style" yaml:"forcePathStyle" mapstructure:"DRIFTHOOK_STORAGE_S3_FORCE_PATH_STYLE"`
	DisableSSL     bool   `json:"disable_ssl" yaml:"disableSSL" mapstructure:"DRIFTHOOK_STORAGE_S3_DISABLE_SSL"`
}

type StorageConfiguration struct {
	S3 S3Configuration `json:"s3" yaml:"s3" mapstructure:",squash"`
}

type Configuration struct {
	Server   ServerConfiguration   `json:"server" yaml:"server" mapstructure:",squash"`
	Database DatabaseConfiguration `json:"database" yaml:"database" mapstructure:",squash"`
	Queue    QueueConfiguration    `json:"queue" yaml:"queue" mapstructure:",squash"`
	Storage  StorageConfiguration  `json:"storage" yaml:"storage" mapstructure:",squash"`
}

var DefaultConfiguration = Configuration{
	Server: ServerConfiguration{
		Host: "localhost",
		Port: 8100,
		Env:  "dev",
	},

	Database: DatabaseConfiguration{
		Postgres: PostgresConfiguration{
			Host: "localhost",
			Port: 5432,
			User: PostgresUserConfiguration{
				Username: "postgres",
				Password: "postgres",
			},
			DatabaseName:      "drifthook",
			Options:           "sslmode=disable",
			MaxOpenConnection: 15,
			MaxIdleConnection: 12,
		},
	},

	Queue: QueueConfiguration{
		NATS: NATSConfiguration{
			Urls:     []string{"nats://localhost:4222"},
			Username: "nats",
			Password: "nats",
		},
	},
}
