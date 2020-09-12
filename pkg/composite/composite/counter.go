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

// QuackCounter Декоратор для подсчета кряков.
type QuackCounter struct {
	duck    Quackable
	counter *counter
}

// Quack Крякнуть.
func (q *QuackCounter) Quack() {
	q.duck.Quack()
	q.counter.increase()
}

// GetQuacks Получить количество кряков.
func (q *QuackCounter) GetQuacks() int {
	return q.counter.getCount()
}

// NewQuackCounter Создать декоратор.
func NewQuackCounter(duck Quackable) *QuackCounter {
	counter := getCounter()

	return &QuackCounter{duck, counter}
}
