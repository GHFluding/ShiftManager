# README.md

# Shift Management 

A Go-based microservices system for managing shifts, tasks, and Bitrix24 chatbot integration.

## Project Structure

├── SM - Core service with database
├── link - Linking service
├── bitrixSM - Bitrix24 chatbot integration
└── contract
    ├──contract/protos - protobuf documentation
    └── gen/go - Generated gRPC code

## Services

1. **SM (Shift Manager)**
   - Main service with PostgreSQL database
   - REST API endpoints
   - Database migrations

2. **Link Service**
   - Handles service-to-service communication
   - GRPC API endpoints
   - API Gateway functionality

3. **BitrixSM**
   - Bitrix24 chatbot integration
   - Business logic for chat interactions

## Prerequisites

- Go 1.24+
- Docker and Docker Compose
- Protobuf compiler (protoc)
- Go plugins:
  protoc-gen-go v1.36.6
  protoc-gen-go-grpc 1.3.0
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
