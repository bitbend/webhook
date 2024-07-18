package config

type Config struct {
	ServerHost string `viper:"DRIFTHOOK_SERVER_HOST"`
	ServerPort int    `viper:"DRIFTHOOK_SERVER_PORT"`

	DatabaseHost     string `viper:"DRIFTHOOK_DATABASE_HOST"`
	DatabasePort     int    `viper:"DRIFTHOOK_DATABASE_PORT"`
	DatabaseUsername string `viper:"DRIFTHOOK_DATABASE_USERNAME"`
	DatabasePassword string `viper:"DRIFTHOOK_DATABASE_PASSWORD"`
	DatabaseName     string `viper:"DRIFTHOOK_DATABASE_NAME"`
	DatabaseOptions  string `viper:"DRIFTHOOK_DATABASE_OPTIONS"`

	QueueUrls     []string `viper:"DRIFTHOOK_QUEUE_URLS"`
	QueueUsername string   `viper:"DRIFTHOOK_QUEUE_USERNAME"`
	QueuePassword string   `viper:"DRIFTHOOK_QUEUE_PASSWORD"`
}

var DefaultConfig = Config{
	ServerHost: "localhost",
	ServerPort: 8080,

	DatabaseHost:     "localhost",
	DatabasePort:     5432,
	DatabaseUsername: "postgres",
	DatabasePassword: "postgres",
	DatabaseName:     "drifthook",
	DatabaseOptions:  "",

	QueueUrls:     []string{"nats://localhost:4222"},
	QueueUsername: "",
	QueuePassword: "",
}
