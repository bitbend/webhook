package config

type Configuration struct {
	ServerHost string `mapstructure:"DRIFTHOOK_SERVER_HOST"`
	ServerPort uint   `mapstructure:"DRIFTHOOK_SERVER_PORT"`

	PostgresHost         string `mapstructure:"DRIFTHOOK_POSTGRES_HOST"`
	PostgresPort         uint   `mapstructure:"DRIFTHOOK_POSTGRES_PORT"`
	PostgresUserUsername string `mapstructure:"DRIFTHOOK_POSTGRES_USER_USERNAME"`
	PostgresUserPassword string `mapstructure:"DRIFTHOOK_POSTGRES_USER_PASSWORD"`
	PostgresDatabaseName string `mapstructure:"DRIFTHOOK_POSTGRES_DATABASE_NAME"`
	PostgresOptions      string `mapstructure:"DRIFTHOOK_POSTGRES_OPTIONS"`

	NATSUrl      string `mapstructure:"DRIFTHOOK_NATS_URL"`
	NATSUsername string `mapstructure:"DRIFTHOOK_NATS_USERNAME"`
	NATSPassword string `mapstructure:"DRIFTHOOK_NATS_PASSWORD"`
}

var DefaultConfiguration = Configuration{
	ServerHost: "localhost",
	ServerPort: 8080,

	PostgresHost:         "localhost",
	PostgresPort:         5432,
	PostgresUserUsername: "postgres",
	PostgresUserPassword: "postgres",
	PostgresDatabaseName: "drifthook",
	PostgresOptions:      "",

	NATSUrl:      "nats://localhost:4222",
	NATSUsername: "",
	NATSPassword: "",
}
