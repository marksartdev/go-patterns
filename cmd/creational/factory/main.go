package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/Mark-Sart/go-patterns/pkg/creational/factory"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	ch := make(chan string)
	ctx := context.Background()
	ctx, finish := context.WithCancel(ctx)
	wg := new(sync.WaitGroup)

	cnt := 0

	pizzaTypes := [4]string{factory.Cheese, factory.Pepperoni, factory.Clam, factory.Veggie}

	var (
		simpleNYPizzaFactory, simpleChicagoPizzaFactory factory.SimplePizzaFactory
		simpleNYPizzaStore, simpleChicagoPizzaStore     factory.PizzaStore
		nyPizzaStore, chicagoPizzaStore                 factory.PizzaStore
	)

	simpleNYPizzaFactory = factory.NewSimpleNYPizzaFactory()
	simpleNYPizzaStore = factory.NewSimplePizzaStore(simpleNYPizzaFactory)
	simpleChicagoPizzaFactory = factory.NewSimpleChicagoPizzaFactory()
	simpleChicagoPizzaStore = factory.NewSimplePizzaStore(simpleChicagoPizzaFactory)
	nyPizzaStore = factory.NewNYPizzaStore()
	chicagoPizzaStore = factory.NewChicagoPizzaStore()

	go startPizzaStore(ctx, wg, ch, simpleNYPizzaStore, "New-York (simple)")
	cnt++

	go startPizzaStore(ctx, wg, ch, simpleChicagoPizzaStore, "Chicago (simple)")
	cnt++

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
	pizzaStore factory.PizzaStore,
	pizzaStoreName string,
) {
	var pizza factory.Pizza

	for {
		select {
		case <-ctx.Done():
			time.Sleep(time.Second)
			fmt.Printf("%s pizzaStore is closing...\n", pizzaStoreName)
			wg.Done()

			return
		case pizzaType := <-ch:
			pizza = pizzaStore.OrderPizza(pizzaType)
			log := pizza.GetLog()

			buffer := fmt.Sprintf("%s:\n", pizzaStoreName)
			for _, step := range log {
				buffer += fmt.Sprintf("\t\t%s\n", step)
			}

			fmt.Println(buffer)
		}
	}
}
