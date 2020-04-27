package strategy

// Flyer Интерфейс способности летать
type Flyer interface {
	fly() string
}

// FlyNoWay Реализация для уток, которые не умеют летать
type FlyNoWay struct{}

// Полететь.
func (f *FlyNoWay) fly() string {
	return "I can't fly!!"
}

// FlyWithWings Реализация для летающих уток
type FlyWithWings struct{}

// Полететь.
func (f *FlyWithWings) fly() string {
	return "I'm flying!!"
}

// FlyRocketPowered Реализация реактивного полета
type FlyRocketPowered struct{}

// Полететь.
func (f *FlyRocketPowered) fly() string {
	return "I'm flying with a rocket!!"
}
