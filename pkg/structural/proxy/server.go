package proxy

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

// Служба.
type gumballMachineServer struct {
	machine GumballMachine
}

// GetLocation Получить местоположение автомата.
func (g gumballMachineServer) GetLocation(context.Context, *Nothing) (*Location, error) {
	response := &Location{}
	response.Data = g.machine.GetLocation()

	return response, nil
}

// GetCount Получить количество оставшихся шариков.
func (g gumballMachineServer) GetCount(context.Context, *Nothing) (*Count, error) {
	response := &Count{}
	response.Data = int32(g.machine.GetCount())

	return response, nil
}

// GetState Получить текущее состояние.
func (g gumballMachineServer) GetState(context.Context, *Nothing) (*State, error) {
	response := &State{}
	response.Data = fmt.Sprint(g.machine.GetState())

	return response, nil
}

// StartService Запустить службу.
func StartService(ctx context.Context, wg *sync.WaitGroup, machine GumballMachine, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterGumballMachineRemoteServer(grpcServer, gumballMachineServer{machine})

	log.Printf("Starting server at %s\n", addr)

	go startServer(wg, grpcServer, lis)
	go stopServer(ctx, grpcServer)

	return nil
}

// Запустить сервер.
func startServer(wg *sync.WaitGroup, grpcServer *grpc.Server, lis net.Listener) {
	defer wg.Done()

	// nolint:gocritic // If catch err => exit(1)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

// Остановить сервер.
func stopServer(ctx context.Context, grpcServer *grpc.Server) {
	<-ctx.Done()

	log.Println("Stopping server...")

	grpcServer.Stop()
}
