package factory

// SimplePizzaProperties Свойства пиццы
type SimplePizzaProperties struct {
	Name       string
	Style      string
	IsPrepared bool
	IsBaked    bool
	IsCutted   bool
	IsBoxed    bool
}

// SimplePizza Интерфейс пиццы
type SimplePizza interface {
	prepare()
	bake()
	cut()
	box()
	GetProperties() SimplePizzaProperties
}

// Базовая структура пиццы
type simplePizza struct {
	SimplePizzaProperties
}

// Приготовить пиццу
func (s *simplePizza) prepare() {
	s.IsPrepared = true
}

// Испечь пиццу
func (s *simplePizza) bake() {
	s.IsBaked = true
}

// Разрезать пиццу
func (s *simplePizza) cut() {
	s.IsCutted = true
}

// Упаковать пиццу
func (s *simplePizza) box() {
	s.IsBoxed = true
}

// GetProperties Получить свойства пиццы
func (s *simplePizza) GetProperties() SimplePizzaProperties {
	return s.SimplePizzaProperties
}
