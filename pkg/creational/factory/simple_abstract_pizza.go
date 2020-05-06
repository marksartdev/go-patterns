package factory

const (
	// Cheese Сырная пицца
	Cheese = "cheese"
	// Pepperoni Пицца "Пепперони"
	Pepperoni = "pepperoni"
	// Clam Пицца с мидиями
	Clam = "clam"
	// Veggie Вегетарианская пицца
	Veggie = "veggie"
)

// SimplePizzaProperties Свойства пиццы.
type SimplePizzaProperties struct {
	Name       string
	Style      string
	IsPrepared bool
	IsBaked    bool
	IsCutted   bool
	IsBoxed    bool
}

// SimplePizza Интерфейс пиццы.
type SimplePizza interface {
	prepare()
	bake()
	cut()
	box()
	GetProperties() SimplePizzaProperties
}

// Абстрактная пицца.
type simpleAbstractPizza struct {
	SimplePizzaProperties
}

// Приготовить пиццу.
func (s *simpleAbstractPizza) prepare() {
	s.IsPrepared = true
}

// Испечь пиццу.
func (s *simpleAbstractPizza) bake() {
	s.IsBaked = true
}

// Разрезать пиццу.
func (s *simpleAbstractPizza) cut() {
	s.IsCutted = true
}

// Упаковать пиццу.
func (s *simpleAbstractPizza) box() {
	s.IsBoxed = true
}

// GetProperties Получить свойства пиццы.
func (s *simpleAbstractPizza) GetProperties() SimplePizzaProperties {
	return s.SimplePizzaProperties
}
