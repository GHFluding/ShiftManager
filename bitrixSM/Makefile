# Makefile

.PHONY: local dev

local:
	go run ./cmd/main.go

dev:
	docker build -t bitrix .
	docker run -d -p 8080:80 --name	bitrix