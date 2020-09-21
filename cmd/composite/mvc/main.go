package main

import "github.com/marksartdev/go-patterns/pkg/composite/mvc"

func main() {
	// nolint:gocritic // First example.
	// model := mvc.NewBeatModel()
	// controller := mvc.NewBeatController(model)
	//
	// controller.Run()
	//
	heartModel := mvc.NewHeartModel()
	heartController := mvc.NewHeartController(heartModel)

	heartController.Run()
}
