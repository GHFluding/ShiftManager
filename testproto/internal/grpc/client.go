package grpc

import (
	"context"
	"log"
	entities "testproto/internal/gen/entities"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

func RunClient() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := entities.NewMachineServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Create(ctx, &entities.CreateMachine{
		Name:             "Test Machine",
		IsRepairRequired: proto.Bool(true),
		IsActive:         proto.Bool(true),
	})
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	log.Printf("Response: %v", resp.GetData().GetName())
}
