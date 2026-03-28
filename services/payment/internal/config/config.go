package config

type Config struct {
	PostgresDSN  string   `env:"POSTGRES_DSN" env-required:"true"`
	RedisAddr    string   `env:"REDIS_ADDR" env-default:"redis:6379"`
	KafkaBrokers []string `env:"KAFKA_BROKERS" env-required:"true"`
	LogLevel     string   `env:"LOG_LEVEL" env-default:"info"`
}
