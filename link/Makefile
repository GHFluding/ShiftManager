# Makefile

.PHONY: local docker-build

local:
	go run ./cmd/main.go

docker-build:
	docker build -t linksm .
	docker run -d -p 8080:80 --name	linksm