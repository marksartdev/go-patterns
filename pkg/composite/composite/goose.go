package composite

// Гусь.
type goose struct {
	customWriter
}

// Honk Крикнуть.
func (g goose) Honk() {
	g.write("Honk")
}

// Адаптер для гуся.
type gooseAdapter struct {
	goose goose
}

// Quack Крякнуть.
func (g gooseAdapter) Quack() {
	g.goose.Honk()
}
