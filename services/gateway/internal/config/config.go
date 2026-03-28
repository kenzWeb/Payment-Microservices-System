package config

type Config struct {
	HTTPPort    int    `env:"HTTP_PORT" env-default:"8080"`
	AuthAddr    string `env:"AUTH_ADDR" env-default:"auth:50051"`
	OrderAddr   string `env:"ORDER_ADDR" env-default:"order:50051"`
	CatalogAddr string `env:"CATALOG_ADDR" env-default:"catalog:50051"`
	LogLevel    string `env:"LOG_LEVEL" env-default:"info"`
}
