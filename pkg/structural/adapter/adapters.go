package adapter

import (
	"math/rand"
	"strings"
)

const flyConverter = 3

// Утиный адаптер для индюшки.
type turkeyAdapter struct {
	turkey Turkey
}

// Quack Крякает.
func (t turkeyAdapter) Quack() string {
	return t.turkey.Gobble()
}

// Fly Летает.
func (t turkeyAdapter) Fly() string {
	log := make([]string, flyConverter)

	for i := 0; i < flyConverter; i++ {
		log[i] = t.turkey.Fly()
	}

	return strings.Join(log, "\n")
}

// NewTurkeyAdapter Создает утиный адаптер для индюшки.
func NewTurkeyAdapter(turkey Turkey) Duck {
	return turkeyAdapter{turkey}
}

// Индюшачий адаптер для утки.
type duckAdapter struct {
	duck   Duck
	random *rand.Rand
}

// Gobble Кулдыкает.
func (d duckAdapter) Gobble() string {
	return d.duck.Quack()
}

// Fly Летает.
func (d duckAdapter) Fly() string {
	if d.random.Intn(flyConverter) == 0 {
		return d.duck.Fly()
	}

	return ""
}

// NewDuckAdapter Создает индюшачий адаптер для утки.
func NewDuckAdapter(duck Duck, seed int64) Turkey {
	source := rand.NewSource(seed)
	random := rand.New(source)

	return duckAdapter{duck, random}
}
