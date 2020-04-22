package factory

// Интерфейс методов получения информации о пицце
type getters interface {
	GetName() string
	IsPrepared() bool
	IsBaked() bool
	IsCutted() bool
	IsBoxed() bool
}

// Pizza Интерфейс пиццы
type Pizza interface {
	getters
	prepare()
	bake()
	cut()
	box()
}

// Базовая структура пиццы
type basePizza struct {
	name       string
	isPrepared bool
	isBaked    bool
	isCutted   bool
	isBoxed    bool
}

// GetName Получить название пиццы
func (b *basePizza) GetName() string {
	return b.name
}

// IsPrepared Получить статус подготовки пиццы
func (b *basePizza) IsPrepared() bool {
	return b.isPrepared
}

// IsBaked Получить статус запекания пиццы
func (b *basePizza) IsBaked() bool {
	return b.isBaked
}

// IsCutted Получить статус разрезания пиццы
func (b *basePizza) IsCutted() bool {
	return b.isCutted
}

// IsBoxed Получить статус упаковки пиццы
func (b *basePizza) IsBoxed() bool {
	return b.isBoxed
}

// Приготовить пиццу
func (b *basePizza) prepare() {
	b.isPrepared = true
}

// Испечь пиццу
func (b *basePizza) bake() {
	b.isBaked = true
}

// Разрезать пиццу
func (b *basePizza) cut() {
	b.isCutted = true
}

// Упаковать пиццу
func (b *basePizza) box() {
	b.isBoxed = true
}

// Сырная пицца
type cheesePizza struct {
	basePizza
}

// Создать сырную пиццу
func newCheesePizza() Pizza {
	pizza := new(cheesePizza)
	pizza.name = "Сырная"

	return pizza
}

// Пицца "Пепперони"
type pepperoniPizza struct {
	basePizza
}

// Создать пиццу "Пепперони"
func newPepperoniPizza() Pizza {
	pizza := new(pepperoniPizza)
	pizza.name = "Пепперони"

	return pizza
}

// Пицца с мидиями
type clamPizza struct {
	basePizza
}

// Создать пиццу с мидиями
func newClamPizza() Pizza {
	pizza := new(clamPizza)
	pizza.name = "С мидиями"

	return pizza
}

// Вегетарианская пицца
type veggiePizza struct {
	basePizza
}

// Создать вегетарианскую пиццу
func newVeggiePizza() Pizza {
	pizza := new(veggiePizza)
	pizza.name = "Вегетарианская"

	return pizza
}
