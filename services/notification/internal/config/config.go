package config

type Config struct {
	KafkaBrokers []string `env:"KAFKA_BROKERS" env-required:"true"`
	LogLevel     string   `env:"LOG_LEVEL" env-default:"info"`
}
