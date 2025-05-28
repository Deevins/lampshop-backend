OUT_DIR := gen

# находим все .proto файлы
PROTO_FILES := $(shell find api -type f -name "*.proto")
# формируем список целевых .pb.go
GEN_FILES := $(patsubst api/%, $(OUT_DIR)/%, $(PROTO_FILES:.proto=.pb.go))

.PHONY: all protos clean
all: protos

# собираем gen и файлы
protos: | ensure-out-dir $(GEN_FILES)

# создаём корень OUT_DIR без ошибок
ensure-out-dir:
	@mkdir -p $(OUT_DIR)

# основное правило: прямое соответствие api/.../X.proto -> gen/.../X.pb.go
$(OUT_DIR)/%.pb.go: api/%.proto
	@mkdir -p $(dir $@)
	@echo "[proto] Generating Go code from $<"
	protoc \
		--proto_path=api \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$<

clean:
	rm -rf $(OUT_DIR)