package client

import (
	"time"

	"google.golang.org/grpc"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
)

type Client struct {
	conn    *grpc.ClientConn
	Clients ClientsServices
}

type ClientsServices struct {
	Machine entities.MachineServiceClient
	User    entities.UserServiceClient
	Task    entities.TaskServiceClient
	Shift   entities.ShiftServiceClient
}

func New(target string) (*Client, error) {
	conn, err := grpc.NewClient(target)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		Clients: ClientsServices{
			Machine: entities.NewMachineServiceClient(conn),
			User:    entities.NewUserServiceClient(conn),
			Task:    entities.NewTaskServiceClient(conn),
			Shift:   entities.NewShiftServiceClient(conn),
		},
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

const timeout = 3 * time.Second
