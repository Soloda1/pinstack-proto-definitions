PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc
PROTO_DIR := proto
GO_OUT := gen/go

PROTO_FILES := \
	$(PROTO_DIR)/user/user.proto \
	$(PROTO_DIR)/auth/auth.proto \
	$(PROTO_DIR)/post/post.proto \
	$(PROTO_DIR)/relation/relation.proto \

generate:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--proto_path=include \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		$(PROTO_FILES)
