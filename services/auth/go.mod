module github.com/user/payment-microservices/services/auth

go 1.26.1

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/user/payment-microservices/api/proto v0.0.0-00010101000000-000000000000
	github.com/user/payment-microservices/pkg/config v0.0.0
	github.com/user/payment-microservices/pkg/grpcutil v0.0.0
	github.com/user/payment-microservices/pkg/logger v0.0.0
	google.golang.org/grpc v1.80.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/net v0.52.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260120221211-b8f7ae30c516 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace github.com/user/payment-microservices/pkg/config => ../../pkg/config

replace github.com/user/payment-microservices/pkg/logger => ../../pkg/logger

replace github.com/user/payment-microservices/pkg/postgres => ../../pkg/postgres

replace github.com/user/payment-microservices/pkg/grpcutil => ../../pkg/grpcutil

replace github.com/user/payment-microservices/api/proto => ../../api/proto
