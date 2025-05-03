PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc
PROTO_DIR := proto
GO_OUT := gen/go

generate:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--proto_path=include \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		$(PROTO_DIR)/user/user.proto