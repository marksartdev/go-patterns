package iterator

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// Waitress Интерфейс официантки.
type Waitress interface {
	PrintMenu()
	SetWriter(writer io.Writer)
}

// Официантка.
type waitress struct {
	pancakeHouseMenu Menu
	dinerMenu        Menu
	writer           io.Writer
}

// PrintMenu Печатает меню.
func (w *waitress) PrintMenu() {
	pancakeIterator := w.pancakeHouseMenu.CreateIterator()
	dinerIterator := w.dinerMenu.CreateIterator()
	w.write("MENU\n----\nBREAKFAST")
	w.printMenu(pancakeIterator)
	w.write("\nLUNCH")
	w.printMenu(dinerIterator)
}

// SetWriter Устанавливает writer.
func (w *waitress) SetWriter(writer io.Writer) {
	w.writer = writer
}

// Печатает меню из итератора.
func (w *waitress) printMenu(iterator common.Iterator) {
	for iterator.HasNext() {
		if item, ok := iterator.Next().(MenuItem); ok {
			msg := fmt.Sprintf("%s, %0.2f -- %s", item.GetName(), item.GetPrice(), item.GetDescription())
			w.write(msg)
		}
	}
}

// Пишет в writer.
func (w *waitress) write(msg string) {
	_, err := fmt.Fprintln(w.writer, msg)
	if err != nil {
		log.Println(err)
	}
}

// NewWaitress Создает официантку.
func NewWaitress(pancakeHouseMenu, dinerMenu Menu) Waitress {
	return &waitress{pancakeHouseMenu, dinerMenu, os.Stdout}
}
