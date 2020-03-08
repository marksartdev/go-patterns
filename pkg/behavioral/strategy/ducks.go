package strategy

// Интерфейс отображения
type displayer interface {
	Display() string
}

// Интерфейс плавания
type swimmer interface {
	Swim() string
}

// Интерфейс кряканья
type quackPerformer interface {
	PerformQuack(int) string
	SetQuacker(Quacker)
}

// Интерфейс полетов
type flyPerformer interface {
	PerformFly() string
	SetFlyer(Flyer)
}

// Duck Интерфейс утки
type Duck interface {
	displayer
	swimmer
	quackPerformer
	flyPerformer
}

/*
Базовая структура утки
quacker - Алгоритм кряканья
flyer - Алгоритм полетов
*/
type baseDuck struct {
	quacker Quacker
	flyer   Flyer
}

// Swim Плыть
func (d *baseDuck) Swim() string {
	return "All ducks float, even decoys!"
}

// PerformQuack Крякнуть
func (d *baseDuck) PerformQuack(count int) string {
	return d.quacker.quack(count)
}

// PerformFly Полететь
func (d *baseDuck) PerformFly() string {
	return d.flyer.fly()
}

// SetQuacker Установить алгоритм кряканья
func (d *baseDuck) SetQuacker(quacker Quacker) {
	d.quacker = quacker
}

// SetFlyer Установить алгоритм полетов
func (d *baseDuck) SetFlyer(flyer Flyer) {
	d.flyer = flyer
}

// Кряква
type mallardDuck struct {
	baseDuck
}

// Display Внешний вид кряквы
func (d *mallardDuck) Display() string {
	return "I'm a mallard duck"
}

// NewMallardDuck Создать крякву
func NewMallardDuck() Duck {
	duck := new(mallardDuck)
	duck.quacker = new(Quack)
	duck.flyer = new(FlyWithWings)

	return duck
}

// Красноголовый нырок
type redheadDuck struct {
	baseDuck
}

// Display Внешний вид красноголового нырка
func (d *redheadDuck) Display() string {
	return "I'm a redhead duck"
}

// NewRedheadDuck Создать красноголового нырка
func NewRedheadDuck() Duck {
	duck := new(redheadDuck)
	duck.quacker = new(Quack)
	duck.flyer = new(FlyWithWings)

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
func NewRubberDuck() Duck {
	duck := new(rubberDuck)
	duck.quacker = new(Squeak)
	duck.flyer = new(FlyNoWay)

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
func NewDecoyDuck() Duck {
	duck := new(decoyDuck)
	duck.quacker = new(MuteQuack)
	duck.flyer = new(FlyNoWay)

	return duck
}

// Утка-приманка
type modelDuck struct {
	baseDuck
}

// Display Внешний вид утки приманки
func (d *modelDuck) Display() string {
	return "I'm a model duck"
}

// NewModelDuck Создать утку-приманку
func NewModelDuck() Duck {
	duck := new(modelDuck)
	duck.quacker = new(Quack)
	duck.flyer = new(FlyNoWay)

	return duck
}
