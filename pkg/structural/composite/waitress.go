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
	allMenu MenuComponent
	writer  io.Writer
}

// PrintMenu Печатает меню.
func (w *waitress) PrintMenu() error {
	return w.allMenu.Print()
}

// PrintVegetarianMenu Печатает вегетарианское меню.
func (w *waitress) PrintVegetarianMenu() error {
	if _, err := fmt.Fprintln(w.writer, "\nVEGETARIAN MENU\n----"); err != nil {
		return err
	}

	iterator := w.allMenu.CreateIterator()
	for iterator.HasNext() {
		if component, ok := iterator.Next().(MenuComponent); ok {
			if isVegetarian, err := component.IsVegetarian(); err == nil && isVegetarian {
				if err = component.Print(); err != nil {
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
}

// NewWaitress Создает официантку.
func NewWaitress(allMenu MenuComponent) Waitress {
	return &waitress{allMenu, os.Stdout}
}
