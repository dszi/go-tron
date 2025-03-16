# Makefile for generating gRPC .pb.go files

PROTO_DIR = proto
PB_DIR = pb

PROTOC = protoc
PROTOC_GEN_GO = protoc-gen-go

.PHONY: check-tools
check-tools:
	@command -v $(PROTOC) >/dev/null 2>&1 || { echo >&2 "Error: protoc is not installed."; exit 1; }
	@command -v $(PROTOC_GEN_GO) >/dev/null 2>&1 || { echo >&2 "Error: protoc-gen-go is not installed."; exit 1; }

PROTO_FILES := $(shell find $(PROTO_DIR) -name "*.proto")

.PHONY: generate
generate: check-tools
	@mkdir -p $(PB_DIR)
	$(PROTOC) --go_out=$(PB_DIR) --go_opt=paths=source_relative \
	          --go-grpc_out=$(PB_DIR) --go-grpc_opt=paths=source_relative \
	          -I $(PROTO_DIR) -I $(PROTO_DIR)/google/api -I $(PROTO_DIR)/google/protobuf $(PROTO_FILES)

.PHONY: clean
clean:
	rm -rf $(PB_DIR)/*.pb.go