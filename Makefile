OUT_DIR := gen

# находим все .proto файлы
PROTO_FILES := $(shell find api -type f -name "*.proto")
# соответствующие .pb.go файлы в OUT_DIR (с сохранением структуры api/...)
GEN_FILES := $(patsubst api/%, $(OUT_DIR)/%, $(PROTO_FILES:.proto=.pb.go))

.PHONY: all protos clean
all: protos

# собираем gen и файлы
protos: | ensure-out-dir $(GEN_FILES)

# создаём корневую папку OUT_DIR без ошибок
ensure-out-dir:
	@mkdir -p $(OUT_DIR)

# правило генерации: source_relative для относительных путей
$(OUT_DIR)/%.pb.go: api/%.proto
	@mkdir -p $(dir $@)
	@echo "[proto] Generating Go code from $<"
	protoc \
		--proto_path=api \
		--go_out=paths=source_relative:$(OUT_DIR) \
		--go-grpc_out=paths=source_relative:$(OUT_DIR) \
		$<

clean:
	rm -rf $(OUT_DIR)
