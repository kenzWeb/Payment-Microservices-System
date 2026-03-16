module github.com/user/payment-microservices/services/notification

go 1.26.1

require (
	github.com/user/payment-microservices/pkg/config v0.0.0
	github.com/user/payment-microservices/pkg/kafka v0.0.0
	github.com/user/payment-microservices/pkg/logger v0.0.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.18.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.26 // indirect
	github.com/twmb/franz-go v1.21.0 // indirect
	github.com/twmb/franz-go/pkg/kmsg v1.13.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace (
	github.com/user/payment-microservices/pkg/config => ../../pkg/config
	github.com/user/payment-microservices/pkg/kafka => ../../pkg/kafka
	github.com/user/payment-microservices/pkg/logger => ../../pkg/logger
)
