package config

type ServerConfiguration struct {
	ServerHost        string `mapstructure:"DRIFTHOOK_SERVER_HOST"`
	ServerPort        uint   `mapstructure:"DRIFTHOOK_SERVER_PORT"`
	ServerEnvironment string `mapstructure:"DRIFTHOOK_SERVER_ENVIRONMENT"`
}

type DatabaseConfiguration struct {
	PostgresHost         string `mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_HOST"`
	PostgresPort         uint   `mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_PORT"`
	PostgresUserUsername string `mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_USER_USERNAME"`
	PostgresUserPassword string `mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_USER_PASSWORD"`
	PostgresDatabaseName string `mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_DATABASE_NAME"`
	PostgresOptions      string `mapstructure:"DRIFTHOOK_DATABASE_POSTGRES_OPTIONS"`
}

type QueueConfiguration struct {
	NATSUrl      string `mapstructure:"DRIFTHOOK_QUEUE_NATS_URL"`
	NATSUsername string `mapstructure:"DRIFTHOOK_QUEUE_NATS_USERNAME"`
	NATSPassword string `mapstructure:"DRIFTHOOK_QUEUE_NATS_PASSWORD"`
}

type StorageConfiguration struct {
	S3AccessKey      string `mapstructure:"DRIFTHOOK_STORAGE_S3_ACCESS_KEY"`
	S3SecretKey      string `mapstructure:"DRIFTHOOK_STORAGE_S3_SECRET_KEY"`
	S3BucketName     string `mapstructure:"DRIFTHOOK_STORAGE_S3_BUCKET_NAME"`
	S3PathPrefix     string `mapstructure:"DRIFTHOOK_STORAGE_S3_PATH_PREFIX"`
	S3Region         string `mapstructure:"DRIFTHOOK_STORAGE_S3_REGION"`
	S3ForcePathStyle bool   `mapstructure:"DRIFTHOOK_STORAGE_S3_FORCE_PATH_STYLE"`
	S3DisableSSL     bool   `mapstructure:"DRIFTHOOK_STORAGE_S3_DISABLE_SSL"`
}

type Configuration struct {
	ServerConfiguration   ServerConfiguration   `mapstructure:",squash"`
	DatabaseConfiguration DatabaseConfiguration `mapstructure:",squash"`
	QueueConfiguration    QueueConfiguration    `mapstructure:",squash"`
	StorageConfiguration  StorageConfiguration  `mapstructure:",squash"`
}

var DefaultConfiguration = Configuration{
	ServerConfiguration: ServerConfiguration{
		ServerHost: "localhost",
		ServerPort: 80,
	},

	DatabaseConfiguration: DatabaseConfiguration{
		PostgresHost:         "localhost",
		PostgresPort:         5432,
		PostgresUserUsername: "postgres",
		PostgresUserPassword: "postgres",
		PostgresDatabaseName: "drifthook",
		PostgresOptions:      "",
	},

	QueueConfiguration: QueueConfiguration{
		NATSUrl:      "nats://localhost:4222",
		NATSUsername: "nats",
		NATSPassword: "nats",
	},
}
