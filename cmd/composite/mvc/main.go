package main

import "github.com/marksartdev/go-patterns/pkg/composite/mvc"

func main() {
	model := mvc.NewBeatModel()
	controller := mvc.NewBeatController(model)

	controller.Run()
}
