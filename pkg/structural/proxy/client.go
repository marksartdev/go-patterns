package proxy

import (
	"context"
	"log"
	"sync"

	"google.golang.org/grpc"
)

// NewGumballMachineStub Создать заместителя.
func NewGumballMachineStub(ctx context.Context, wg *sync.WaitGroup, addr string) (GumballMachineRemoteClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	go closeConn(ctx, wg, conn, addr)

	client := NewGumballMachineRemoteClient(conn)

	log.Printf("Connecting to %s\n", addr)

	return client, nil
}

// Закрыть соединение.
func closeConn(ctx context.Context, wg *sync.WaitGroup, conn *grpc.ClientConn, addr string) {
	<-ctx.Done()

	log.Printf("Closing connection to %s\n", addr)

	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}
