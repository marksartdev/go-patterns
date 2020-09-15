package composite

import (
	"io"
	"sync"
)

// nolint:gochecknoglobals // Singleton
var (
	instance *counter
	once     sync.Once
)

// Счетчик.
type counter struct {
	count int
	rw    sync.RWMutex
}

// Увеличить счетчик.
func (c *counter) increase() {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.count++
}

// Получить текущее количество.
func (c *counter) getCount() int {
	c.rw.RLock()
	defer c.rw.RUnlock()

	count := c.count

	return count
}

// Получить инстанс счетчика.
func getCounter() *counter {
	once.Do(func() {
		instance = &counter{}
	})

	return instance
}

// GetQuacks Получить количество кряков.
func GetQuacks() int {
	counter := getCounter()

	return counter.getCount()
}

// ResetCounter Сброс счетчика.
func ResetCounter() {
	counter := getCounter()

	counter.count = 0
}

// Декоратор для подсчета кряков.
type quackCounter struct {
	duck    Quackable
	counter *counter
}

// Quack Крякнуть.
func (q *quackCounter) Quack() {
	q.duck.Quack()
	q.counter.increase()
}

// RegisterObserver Зарегистрировать наблюдателя.
func (q *quackCounter) RegisterObserver(observer observer) {
	q.duck.RegisterObserver(observer)
}

// Оповестить наблюдателей (заглушка, т.к. это делает утка при кряканье).
func (q *quackCounter) notifyObserver() {}

// SetWriter Установить writer.
func (q *quackCounter) SetWriter(writer io.Writer) {
	q.duck.SetWriter(writer)
}

// Создать декоратор.
func newQuackCounter(duck Quackable) Quackable {
	counter := getCounter()

	return &quackCounter{duck, counter}
}
