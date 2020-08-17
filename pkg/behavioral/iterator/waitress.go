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
	menus  common.ArrayList
	writer io.Writer
}

// PrintMenu Печатает меню.
func (w *waitress) PrintMenu() {
	menuIterator := w.menus.Iterator()
	for menuIterator.HasNext() {
		if menu, ok := menuIterator.Next().(Menu); ok {
			w.printMenu(menu.CreateIterator())
		}
	}
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
func NewWaitress(menus common.ArrayList) Waitress {
	return &waitress{menus, os.Stdout}
}
