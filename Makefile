OUT_DIR := gen

PROTO_FILES := $(shell find api -type f -name "*.proto")

GEN_FILES := $(patsubst api/%, $(OUT_DIR)/%, $(PROTO_FILES:.proto=.pb.go))

.PHONY: all protos clean
all: protos

protos: | ensure-out-dir $(GEN_FILES)

ensure-out-dir:
	@mkdir -p $(OUT_DIR)

$(OUT_DIR)/%/%.pb.go: api/%.proto
	@mkdir -p $(dir $@)
	@echo "[proto] Generating Go code from $<"
	protoc \
		--proto_path=api \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$<

clean:
	rm -rf $(OUT_DIR)