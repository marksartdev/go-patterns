package composite

import "sync"

// nolint:gochecknoglobals // Singleton
var (
	instance *counter
	once     sync.Once
)

// Счетчик.
type counter struct {
	count int
	mu    sync.RWMutex
}

// Увеличить счетчик.
func (c *counter) increase() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

// Получить текущее количество.
func (c *counter) getCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

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

// Создать декоратор.
func newQuackCounter(duck Quackable) Quackable {
	counter := getCounter()

	return &quackCounter{duck, counter}
}
