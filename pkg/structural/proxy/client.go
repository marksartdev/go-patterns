package proxy

import (
	"fmt"

	"google.golang.org/grpc"
)

// NewGumballMachineStub Создать заместителя.
func NewGumballMachineStub(address string) (GumballMachineRemoteClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := NewGumballMachineRemoteClient(conn)

	fmt.Printf("Connecting to %s\n", address)

	return client, nil
}
