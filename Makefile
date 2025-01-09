generate-proto:
	cd proto && buf dep update && buf generate --clean && buf generate

PROTOC := protoc
PROTO_DIR := core
OUT_DIR := core
PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)
PROTOC_GEN_SWAGGER := $(shell which protoc-gen-openapiv2)

.PHONY: all
all: generate-protos generate-swagger

.PHONY: generate-protos
generate-protos:
	@echo "Generating .pb.go and .pb.gw.go files..."
	cd core/proto && buf dep update && buf generate --clean && buf generate

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	find $(OUT_DIR) -type f \( -name "*.pb.go" -o -name "*.pb.gw.go" -o -name "*.swagger.json" \) -delete
