package main

import (
	"time"

	"github.com/marksartdev/go-patterns/pkg/structural/vproxy"
)

func main() {
	ex := vproxy.NewExchanger()
	ex.Show()

	// nolint:gomnd // Для теста
	time.Sleep(3 * time.Second)
}
