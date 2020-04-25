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

	chNY := make(chan string)
	chChicago := make(chan string)
	ctx := context.Background()
	ctx, finish := context.WithCancel(ctx)
	wg := new(sync.WaitGroup)

	cnt := 0

	pizzaTypes := [4]string{"cheese", "pepperoni", "clam", "veggie"}

	go startSimplePizzaStore(ctx, wg, chNY, "New-York")
	cnt++

	go startSimplePizzaStore(ctx, wg, chChicago, "Chicago")
	cnt++

	wg.Add(cnt)

	for i := 0; i < 10; i++ {
		pizzaType := pizzaTypes[random.Intn(len(pizzaTypes))]

		if i%2 == 0 {
			chNY <- pizzaType
		} else {
			chChicago <- pizzaType
		}
	}

	finish()
	wg.Wait()
	fmt.Println("Все пиццерии успешно закрылись!")
}

func startSimplePizzaStore(ctx context.Context, wg *sync.WaitGroup, chOrder chan string, style string) {
	var (
		simpleStore     factory.SimplePizzaStore
		simpleFactory   factory.SimplePizzaFactory
		pizza           factory.SimplePizza
		pizzaProperties factory.SimplePizzaProperties
	)

	if style == "New-York" {
		simpleFactory = factory.NewSimpleNYPizzaFactory()
	} else {
		simpleFactory = factory.NewSimpleChicagoPizzaFactory()
	}

	simpleStore = factory.NewSimplePizzaStore(simpleFactory)

	for {
		select {
		case <-ctx.Done():
			time.Sleep(time.Second)
			fmt.Printf("Пиццерия %s закрывается...\n", style)
			wg.Done()

			return
		case pizzaType := <-chOrder:
			pizza = simpleStore.OrderPizza(pizzaType)
			pizzaProperties = pizza.GetProperties()

			fmt.Printf(
				"%-12s Название: %-18s Приготовлена: %-8v Выпечена: %-8v Разрезана: %-8v Упакована: %-8v\n",
				pizzaProperties.Style,
				pizzaProperties.Name,
				pizzaProperties.IsPrepared,
				pizzaProperties.IsBaked,
				pizzaProperties.IsCutted,
				pizzaProperties.IsBoxed,
			)
		}
	}
}
