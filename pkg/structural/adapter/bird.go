package adapter

// Duck Интерфейс утки.
type Duck interface {
	Quack() string
	Fly() string
}

// MallardDuck Кряква.
type MallardDuck struct{}

// Quack Крякает.
func (MallardDuck) Quack() string {
	return "Quack"
}

// Fly Летает.
func (MallardDuck) Fly() string {
	return "I'm flying"
}

// Turkey Интерфейс индюшки.
type Turkey interface {
	Gobble() string
	Fly() string
}

// WildTurkey Дикая индюшка.
type WildTurkey struct{}

// Gobble Кулдыкает.
func (WildTurkey) Gobble() string {
	return "Gobble gobble"
}

// Fly Летает.
func (WildTurkey) Fly() string {
	return "I'm flying a short distance"
}
