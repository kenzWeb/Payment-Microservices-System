package config

type Config struct {
	GRPCPort      int      `env:"GRPC_PORT" env-default:"50053"`
	PostgresDSN   string   `env:"POSTGRES_DSN" env-required:"true"`
	InventoryAddr string   `env:"INVENTORY_ADDR" env-default:"inventory:50052"`
	KafkaBrokers  []string `env:"KAFKA_BROKERS" env-required:"true"`
	LogLevel      string   `env:"LOG_LEVEL" env-default:"info"`
}
