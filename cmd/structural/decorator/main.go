package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/marksartdev/go-patterns/pkg/structural/decorator"
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

	sizes := [3]int{decorator.Small, decorator.Medium, decorator.Large}

	source := rand.NewSource(time.Now().UnixNano())
	// nolint:gosec
	random := rand.New(source)

	for i := 0; i < 10; i++ {
		beverage := beverages[random.Intn(len(beverages))]()
		beverage = condiments[random.Intn(len(condiments))](beverage)
		beverage = condiments[random.Intn(len(condiments))](beverage)
		beverage.SetSize(sizes[random.Intn(len(sizes))])

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
