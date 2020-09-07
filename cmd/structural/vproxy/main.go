package main

import (
	"fmt"
	"time"

	"github.com/marksartdev/go-patterns/pkg/structural/vproxy"
)

func main() {
	ex := vproxy.NewExchanger()
	ex.Show()

	fmt.Println(ex.GetRates())

	// nolint:gomnd // Для теста
	time.Sleep(3 * time.Second)

	fmt.Println(ex.GetRates())
}
