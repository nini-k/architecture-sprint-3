PROTO_GEN_PATH=./internal/pb
SWAGGER_UI_VERSION=v5.17.14
SWAGGER_UI_GEN_DIR=./swagger-ui

generate-swaggerui:
	GEN_DIR=$(SWAGGER_UI_GEN_DIR) SWAGGER_UI_VERSION=$(SWAGGER_UI_VERSION) ./scripts/generate-swagger-ui.sh

generate-proto:
	protoc -I ./proto \
    	--go_out ${PROTO_GEN_PATH} \
		--go_opt paths=source_relative \
    	--go-grpc_out ${PROTO_GEN_PATH} \
		--go-grpc_opt paths=source_relative \
		--grpc-gateway_out ${PROTO_GEN_PATH} \
    	--grpc-gateway_opt paths=source_relative \
    	--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out ${SWAGGER_UI_GEN_DIR} \
    	./proto/telemetry/telemetry.proto

local-run:
	MONGO_URL="" GRPC_ENDPOINT=0.0.0.0:9090 GRPC_GATEWAY_ENDPOINT=0.0.0.0:9091 go run cmd/telemetry/main.go
