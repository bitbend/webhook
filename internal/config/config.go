package config

type DatabaseType string

func (dt DatabaseType) String() string {
	return string(dt)
}

const (
	DatabaseTypePostgres DatabaseType = "postgres"
)

type QueueType string

func (qt QueueType) String() string {
	return string(qt)
}

const (
	QueueTypeNATS QueueType = "nats"
)

type Config struct {
	ServerHost string `viper:"DRIFTHOOK_SERVER_HOST"`
	ServerPort int    `viper:"DRIFTHOOK_SERVER_PORT"`

	DatabaseType     string `viper:"DRIFTHOOK_DATABASE_TYPE"`
	DatabaseHost     string `viper:"DRIFTHOOK_DATABASE_HOST"`
	DatabasePort     int    `viper:"DRIFTHOOK_DATABASE_PORT"`
	DatabaseUser     string `viper:"DRIFTHOOK_DATABASE_USER"`
	DatabasePassword string `viper:"DRIFTHOOK_DATABASE_PASSWORD"`
	DatabaseName     string `viper:"DRIFTHOOK_DATABASE_NAME"`
	DatabaseOptions  string `viper:"DRIFTHOOK_DATABASE_OPTIONS"`

	QueueType     string   `viper:"DRIFTHOOK_QUEUE_TYPE"`
	QueueUrls     []string `viper:"DRIFTHOOK_QUEUE_URLS"`
	QueueUser     string   `viper:"DRIFTHOOK_QUEUE_USER"`
	QueuePassword string   `viper:"DRIFTHOOK_QUEUE_PASSWORD"`
}

var DefaultConfig = Config{
	ServerHost: "localhost",
	ServerPort: 8080,

	DatabaseType:     "postgres",
	DatabaseHost:     "localhost",
	DatabasePort:     5432,
	DatabaseUser:     "postgres",
	DatabasePassword: "postgres",
	DatabaseName:     "drifthook",
	DatabaseOptions:  "",

	QueueType:     "nats",
	QueueUrls:     []string{"nats://localhost:4222"},
	QueueUser:     "",
	QueuePassword: "",
}
