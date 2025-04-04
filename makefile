# Makefile

.PHONY: generate build docker-build docker-up 

PROTO_DIR=contract\contract\protos\entities
GEN_DIR=contract\gen\go

generate:
	mkdir -p $(GEN_DIR)
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GEN_DIR) \
		--go-grpc_out=$(GEN_DIR) \
		$(PROTO_DIR)/machine/machine.proto \
		$(PROTO_DIR)/shift/shift.proto \
		$(PROTO_DIR)/task/task.proto \
		$(PROTO_DIR)/user/user.proto

build:
	cd SM && go build -o ../bin/SM ./cmd/sm
	cd link && go build -o ../bin/link ./cmd
	cd bitrixSM && go build -o ../bin/bitrixSM ./cmd

docker-build:
	docker build -t sm-core -f SM/build/Dockerfile .
	docker build -t link-service -f link/Dockerfile .
	docker build -t bitrix-bot -f bitrixSM/Dockerfile .

docker-up:
	docker-compose -f SM/deployments/docker-compose/docker-compose.yaml up -d