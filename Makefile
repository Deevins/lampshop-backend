# Makefile for gRPC code generation
# Генерирует .pb.go в структуре, заданной go_package в .proto

OUT_DIR := gen
PROTO_DIR := api
PROTO_FILES := $(shell find $(PROTO_DIR) -type f -name "*.proto")

.PHONY: all protos clean

# По умолчанию генерируем все
all: protos

# Основной таргет: создаём OUT_DIR и запускаем protoc на всех .proto
protos: | ensure-out-dir
	@echo "[proto] Generating Go code for all .proto files..."
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$(PROTO_FILES)

# Создаёт каталог OUT_DIR, если его нет
ensure-out-dir:
	@mkdir -p $(OUT_DIR)

# Удаляет сгенерированный код
clean:
	@rm -rf $(OUT_DIR)
