module github.com/user/payment-microservices/services/payment

go 1.26.1

require (
	github.com/jackc/pgx/v5 v5.7.1
	github.com/redis/go-redis/v9 v9.7.0
	github.com/user/payment-microservices/pkg/config v0.0.0
	github.com/user/payment-microservices/pkg/kafka v0.0.0
	github.com/user/payment-microservices/pkg/logger v0.0.0
	github.com/user/payment-microservices/pkg/postgres v0.0.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.18.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.26 // indirect
	github.com/twmb/franz-go v1.21.0 // indirect
	github.com/twmb/franz-go/pkg/kmsg v1.13.1 // indirect
	golang.org/x/crypto v0.50.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace (
	github.com/user/payment-microservices/api/proto => ../../api/proto
	github.com/user/payment-microservices/pkg/config => ../../pkg/config
	github.com/user/payment-microservices/pkg/kafka => ../../pkg/kafka
	github.com/user/payment-microservices/pkg/logger => ../../pkg/logger
	github.com/user/payment-microservices/pkg/postgres => ../../pkg/postgres
)
