package main

import "testproto/internal/grpc"

func main() {
	grpc.RunServer()
	grpc.RunClient()
}
