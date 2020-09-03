package proxy

import (
	"context"
	"fmt"
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
func StartService(wg *sync.WaitGroup, machine GumballMachine, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterGumballMachineRemoteServer(grpcServer, gumballMachineServer{machine})

	fmt.Printf("Starting server at %s\n", address)
	wg.Done()

	return grpcServer.Serve(lis)
}
