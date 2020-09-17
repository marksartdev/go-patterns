package main

import "github.com/marksartdev/go-patterns/pkg/composite/mvc"

func main() {
	djView := mvc.NewDJView()
	djView.Run()
}
