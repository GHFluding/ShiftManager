# Makefile

.PHONY: local docker-build

local:
	go run ./cmd/main.go

docker-build:
	docker build -t bitrix .
	docker run -d -p 8080:80 --name	bitrix