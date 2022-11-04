.PHONY: gen help

PROTO_DIR=api/protobuf-spec

help:
	@echo "make gen - 生成pb及grpc代码"

gen:
	protoc \
	--proto_path=api/protobuf-spec \
	--go_out=$(PROTO_DIR) --go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_DIR) --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=$(PROTO_DIR) --grpc-gateway_opt=paths=source_relative \
	$(shell find $(PROTO_DIR) -iname "*.proto")
