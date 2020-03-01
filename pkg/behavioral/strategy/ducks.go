package strategy

// Duck Интерфейс утки
type Duck interface {
	Display() string
	Swim() string
	PerformQuack(count int) string
	PerformFly() string
	SetQuackBehavior(quackBehavior QuackBehavior)
	SetFlyBehavior(flyBehavior FlyBehavior)
}

/*
Базовая структура утки
quackBehavior - Алгоритм кряканья
flyBehavior - Алгоритм полетов
*/
type baseDuck struct {
	quackBehavior QuackBehavior
	flyBehavior   FlyBehavior
}

// Swim Плыть
func (d *baseDuck) Swim() string {
	return "All ducks float, even decoys!"
}

// PerformQuack Крякнуть
func (d *baseDuck) PerformQuack(count int) string {
	return d.quackBehavior.Quack(count)
}

// PerformFly Полететь
func (d *baseDuck) PerformFly() string {
	return d.flyBehavior.Fly()
}

// SetQuackBehavior Установить алгоритм кряканья
func (d *baseDuck) SetQuackBehavior(quackBehavior QuackBehavior) {
	d.quackBehavior = quackBehavior
}

// SetFlyBehavior Установить алгоритм полетов
func (d *baseDuck) SetFlyBehavior(flyBehavior FlyBehavior) {
	d.flyBehavior = flyBehavior
}

// mallardDuck Кряква
type mallardDuck struct {
	baseDuck
}

// Display Внешний вид кряквы
func (d *mallardDuck) Display() string {
	return "I'm a mallard duck"
}

// NewMallardDuck Создать крякву
func NewMallardDuck() *mallardDuck {
	duck := new(mallardDuck)
	duck.quackBehavior = new(Quack)
	duck.flyBehavior = new(FlyWithWings)

	return duck
}

// redheadDuck Красноголовый нырок
type redheadDuck struct {
	baseDuck
}

// Display Внешний вид красноголового нырка
func (d *redheadDuck) Display() string {
	return "I'm a redhead duck"
}

// NewRedheadDuck Создать красноголового нырка
func NewRedheadDuck() *redheadDuck {
	duck := new(redheadDuck)
	duck.quackBehavior = new(Quack)
	duck.flyBehavior = new(FlyWithWings)

	return duck
}

// Резиновая уточка
type rubberDuck struct {
	baseDuck
}

// Display Внешний вид резиновой уточки
func (d *rubberDuck) Display() string {
	return "I'm a rubber duck"
}

// NewRubberDuck Создать резиновую уточку
func NewRubberDuck() *rubberDuck {
	duck := new(rubberDuck)
	duck.quackBehavior = new(Squeak)
	duck.flyBehavior = new(FlyNoWay)

	return duck
}

// Деревянная утка
type decoyDuck struct {
	baseDuck
}

// Display Внешний вид деревянной утки
func (d *decoyDuck) Display() string {
	return "I'm a decoy duck"
}

// NewDecoyDuck Создать деревянную утку
func NewDecoyDuck() *decoyDuck {
	duck := new(decoyDuck)
	duck.quackBehavior = new(MuteQuack)
	duck.flyBehavior = new(FlyNoWay)

	return duck
}

// modelDuck Утка-приманка
type modelDuck struct {
	baseDuck
}

// Display Внешний вид утки приманки
func (d *modelDuck) Display() string {
	return "I'm a model duck"
}

// NewModelDuck Создать утку-приманку
func NewModelDuck() *modelDuck {
	duck := new(modelDuck)
	duck.quackBehavior = new(Quack)
	duck.flyBehavior = new(FlyNoWay)

	return duck
}
