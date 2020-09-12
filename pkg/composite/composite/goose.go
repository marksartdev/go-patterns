package composite

import "os"

// Goose Гусь.
type Goose struct {
	customWriter
}

// Honk Крикнуть.
func (g Goose) Honk() {
	g.write("Honk")
}

// NewGoose Создать гуся.
func NewGoose() Goose {
	g := Goose{}
	g.SetWriter(os.Stdout)

	return g
}

// Адаптер для гуся.
type gooseAdapter struct {
	goose Goose
}

// Quack Крякнуть.
func (g gooseAdapter) Quack() {
	g.goose.Honk()
}

// NewGooseAdapter Создать гусиный адаптер.
func NewGooseAdapter(goose Goose) Quackable {
	return gooseAdapter{goose}
}
