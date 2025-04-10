# Makefile

.PHONY: build docker-build docker-run

build:
	docker build -t sm-core -f SM/build/Dockerfile .

docker-build:
	docker-compose up --build -f SM/deployments/docker-compose.yaml 

docker-run:
	docker-compose up  -f SM/deployments/docker-compose.yaml 
	