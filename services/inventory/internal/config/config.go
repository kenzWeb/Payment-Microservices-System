package config

type Config struct {
	GRPCPort    int    `env:"GRPC_PORT" env-default:"50052"`
	PostgresDSN string `env:"POSTGRES_DSN" env-required:"true"`
	LogLevel    string `env:"LOG_LEVEL" env-default:"info"`
}
