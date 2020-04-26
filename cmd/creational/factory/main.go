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

	go startPizzaStore(ctx, wg, ch, simpleNYPizzaStore, "Нью-Йоркская пиццерия (простая)")
	cnt++

	go startPizzaStore(ctx, wg, ch, simpleChicagoPizzaStore, "Чикагская пиццерия (простая)")
	cnt++

	go startPizzaStore(ctx, wg, ch, nyPizzaStore, "Нью-Йоркская пиццерия")
	cnt++

	go startPizzaStore(ctx, wg, ch, chicagoPizzaStore, "Чикагская пиццерия")
	cnt++

	wg.Add(cnt)

	for i := 0; i < 20; i++ {
		pizzaType := pizzaTypes[random.Intn(len(pizzaTypes))]

		ch <- pizzaType
	}

	finish()
	wg.Wait()
	fmt.Println("Все пиццерии успешно закрылись!")
}

func startPizzaStore(
	ctx context.Context,
	wg *sync.WaitGroup,
	ch chan string,
	pizzaStore factory.PizzaStore,
	pizzaStoreName string,
) {
	var (
		pizza           factory.SimplePizza
		pizzaProperties factory.SimplePizzaProperties
	)

	for {
		select {
		case <-ctx.Done():
			time.Sleep(time.Second)
			fmt.Println("Пиццерия закрывается...")
			wg.Done()

			return
		case pizzaType := <-ch:
			pizza = pizzaStore.OrderPizza(pizzaType)
			pizzaProperties = pizza.GetProperties()

			fmt.Printf(
				"%-40s Стиль: %-12s Название: %-18s Приготовлена: %-8v Выпечена: %-8v Разрезана: %-8v Упакована: %-8v\n",
				pizzaStoreName,
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
