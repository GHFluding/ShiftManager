package grpc

import (
	"context"
	"log"
	"net"
	entities "testproto/internal/gen/entities"

	"google.golang.org/grpc"
)

type machineServer struct {
	entities.UnimplementedMachineServiceServer
}

func (s *machineServer) Create(ctx context.Context, req *entities.CreateMachine) (*entities.MachineResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &entities.MachineResponse{
		Data: req,
	}, nil
}

func RunServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	entities.RegisterMachineServiceServer(s, &machineServer{})

	log.Printf("Server started at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
