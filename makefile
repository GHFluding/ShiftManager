# Makefile

.PHONY: generate build docker-build docker-up dev

BIN_DIR=bin
PROTO_DIR=contract/contract/protos/entities
GEN_DIR=contract/gen/go

generate:
	protoc \
	--proto_path=$(PROTO_DIR) \
	--go_out=$(GEN_DIR) \
	--go-grpc_out=$(GEN_DIR) \
	$(PROTO_DIR)/entities.proto

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/SM  SM/cmd/sm
	go build -o $(BIN_DIR)/link link/cmd
	go build -o $(BIN_DIR)/bitrixSM bitrixSM/cmd

docker-build:
	docker build -t sm-core -f SM/build/Dockerfile .
	docker build -t link-service -f link/Dockerfile .
	docker build -t bitrix-bot -f bitrixSM/Dockerfile .

docker-up:
	docker-compose up -d

dev:
	docker-compose up --build