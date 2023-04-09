#!/usr/bin/env bash

# add tools.go in project
# package tools
#
# import (
# 	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
# 	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
# 	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
# 	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
# )
# init compile tools
go mod tidy
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# init service DI
go install github.com/google/wire/cmd/wire@latest
API_PROTO_FILES=$(find api -name "*.proto")
CONF_PROTO_FILES=$(find configs -name "*.proto")
# auto generate code from proto files

# generate go grpc grpc-gateway and swagger.json
protoc --proto_path=./configs \
  --proto_path=./third_party \
  --go_out=paths=source_relative:./configs \
  "${CONF_PROTO_FILES}"

# generate go grpc grpc-gateway and swagger.json
protoc --proto_path=./api \
  --proto_path=./third_party \
  --go_out=paths=source_relative:./api \
  --go-grpc_out=paths=source_relative:./api \
  --grpc-gateway_out ./api \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --openapiv2_out ./docs/openapiv2 \
  --openapiv2_opt logtostderr=true \
  "${API_PROTO_FILES}"
