// Package templatemethod Паттерн "Шаблонный метод".
package templatemethod

import (
	"fmt"
	"io"
	"log"
	"os"
)

// CaffeineBeverage Интерфейс кофейного напитка.
type CaffeineBeverage interface {
	PrepareRecipe()
	SetReader(reader io.Reader)
	SetWriter(writer io.Writer)
}

// Абстрактный кофейный напиток.
type caffeineBeverage struct {
	reader         io.Reader
	writer         io.Writer
	brew           func()
	addCondiments  func()
	condimentsHook func() bool
}

// PrepareRecipe Готовит напиток.
func (c *caffeineBeverage) PrepareRecipe() {
	c.boilWater()
	c.brew()
	c.pourInCup()

	if c.condimentsHook() {
		c.addCondiments()
	}
}

// Кипятит воду.
func (c *caffeineBeverage) boilWater() {
	c.write("Boiling water")
}

// Переливает напиток в чашку.
func (c *caffeineBeverage) pourInCup() {
	c.write("Pouring into cup")
}

// Перехватчик, решающий, добавлять ли добавки в напиток.
func (c *caffeineBeverage) customerWantsCondimentsBase() bool {
	return true
}

// Печатает сообщение.
func (c *caffeineBeverage) write(msg string) {
	_, err := fmt.Fprintln(c.writer, msg)
	c.errorHandler(err)
}

// Обрабатывает ошибку.
func (c *caffeineBeverage) errorHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// SetReader Устанавливает reader.
func (c *caffeineBeverage) SetReader(reader io.Reader) {
	c.reader = reader
}

// SetWriter Устанавливает writer.
func (c *caffeineBeverage) SetWriter(writer io.Writer) {
	c.writer = writer
}

// Создает базовую структуру кофейного напитка.
func newCaffeineBeverage() *caffeineBeverage {
	c := &caffeineBeverage{}
	c.reader = os.Stdin
	c.writer = os.Stdout
	c.condimentsHook = c.customerWantsCondimentsBase

	return c
}
