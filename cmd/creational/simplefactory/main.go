package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/marksartdev/go-patterns/pkg/creational/simplefactory"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	// nolint:gosec
	random := rand.New(source)

	ch := make(chan string)
	ctx := context.Background()
	ctx, finish := context.WithCancel(ctx)
	wg := new(sync.WaitGroup)

	cnt := 0

	pizzaTypes := [4]string{
		simplefactory.CheesePizza,
		simplefactory.PepperoniPizza,
		simplefactory.ClamPizza,
		simplefactory.VeggiePizza,
	}

	var (
		nyPizzaFactory, chicagoPizzaFactory simplefactory.PizzaFactory
		nyPizzaStore, chicagoPizzaStore     simplefactory.PizzaStore
	)

	nyPizzaFactory = simplefactory.NewNYPizzaFactory()
	chicagoPizzaFactory = simplefactory.NewChicagoPizzaFactory()
	nyPizzaStore = simplefactory.NewPizzaStore(nyPizzaFactory)
	chicagoPizzaStore = simplefactory.NewPizzaStore(chicagoPizzaFactory)

	go startPizzaStore(ctx, wg, ch, nyPizzaStore, "New-York")
	cnt++

	go startPizzaStore(ctx, wg, ch, chicagoPizzaStore, "Chicago")
	cnt++

	wg.Add(cnt)

	for i := 0; i < 10; i++ {
		pizzaType := pizzaTypes[random.Intn(len(pizzaTypes))]

		ch <- pizzaType
	}

	finish()
	wg.Wait()
	fmt.Println("All pizzaStores were closed!")
}

func startPizzaStore(
	ctx context.Context,
	wg *sync.WaitGroup,
	ch chan string,
	pizzaStore simplefactory.PizzaStore,
	pizzaStoreName string,
) {
	var (
		pizza simplefactory.Pizza
		err   error
	)

	for {
		select {
		case <-ctx.Done():
			time.Sleep(time.Second)
			fmt.Printf("%s pizzaStore is closing...\n", pizzaStoreName)
			wg.Done()

			return
		case pizzaType := <-ch:
			buffer := fmt.Sprintf("%s:\n", pizzaStoreName)

			pizza, err = pizzaStore.OrderPizza(pizzaType)
			if err != nil {
				buffer += fmt.Sprintf("\t\tError: %s\n", err)
				fmt.Println(buffer)

				continue
			}

			log := pizza.GetLog()
			for _, step := range log {
				buffer += fmt.Sprintf("\t\t%s\n", step)
			}

			fmt.Println(buffer)
		}
	}
}
