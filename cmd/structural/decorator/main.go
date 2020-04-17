package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"

	"github.com/Mark-Sart/go-patterns/pkg/structural/decorator"
)

func main() {
	beverages := []func() decorator.Beverage{
		decorator.NewHouseBlend,
		decorator.NewDarkRoast,
		decorator.NewDecaf,
		decorator.NewEspresso,
	}

	condiments := []func(decorator.Beverage) decorator.Beverage{
		decorator.NewMilkDecorator,
		decorator.NewMochaDecorator,
		decorator.NewSoyDecorator,
		decorator.NewWhipDecorator,
	}

	for i := 0; i < 10; i++ {
		beverage := beverages[rand.Intn(len(beverages))]()
		beverage = condiments[rand.Intn(len(condiments))](beverage)
		beverage = condiments[rand.Intn(len(condiments))](beverage)

		fmt.Printf("%s: %.2f\n", beverage.GetDescription(), beverage.Cost())
	}

	fmt.Println()

	textBuffer := bytes.NewBuffer([]byte("I know the Decorator Pattern therefore I RULE!"))
	var reader io.Reader = bufio.NewReader(textBuffer)
	reader = decorator.NewLowCaseReader(reader)

	buffer := make([]byte, 8)

	for {
		l, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Print("\n")
			break
		}

		fmt.Print(string(buffer[:l]))
	}
}
