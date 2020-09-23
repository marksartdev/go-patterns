package main

import (
	"github.com/marksartdev/go-patterns/pkg/composite/model2"
	"github.com/marksartdev/go-patterns/pkg/composite/mvc"
)

func main() {
	model := mvc.NewBeatModel()

	model2.StartServer(model, ":8080")
}
