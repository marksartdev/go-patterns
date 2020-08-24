package composite

import (
	"fmt"
	"io"
	"os"
)

// Waitress Интерфейс официантки.
type Waitress interface {
	PrintMenu() error
	PrintVegetarianMenu() error
	SetWriter(writer io.Writer)
}

// Официантка.
type waitress struct {
	allMenu menuComponent
	writer  io.Writer
}

// PrintMenu Печатает меню.
func (w *waitress) PrintMenu() error {
	return w.allMenu.print()
}

// PrintVegetarianMenu Печатает вегетарианское меню.
func (w *waitress) PrintVegetarianMenu() error {
	if _, err := fmt.Fprintln(w.writer, "\nVEGETARIAN MENU\n----"); err != nil {
		return err
	}

	iterator := w.allMenu.createIterator()
	for iterator.HasNext() {
		if next, ok := iterator.Next().(menuComponent); ok {
			if isVegetarian, err := next.isVegetarian(); err == nil && isVegetarian {
				if err = next.print(); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// SetWriter Устанавливает writer.
func (w *waitress) SetWriter(writer io.Writer) {
	w.writer = writer
	w.allMenu.setWriter(writer)
}

// newWaitress Создает официантку.
func newWaitress(allMenu menuComponent) Waitress {
	return &waitress{allMenu, os.Stdout}
}
