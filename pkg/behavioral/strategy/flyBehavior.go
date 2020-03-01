package strategy

// FlyBehavior Интерфейс способности летать
type FlyBehavior interface {
	Fly() string
}

// FlyNoWay Реализация для уток, которые не умеют летать
type FlyNoWay struct{}

// Fly Полететь
func (f *FlyNoWay) Fly() string {
	return "I can't fly!!"
}

// FlyWithWings Реализация для летающих уток
type FlyWithWings struct{}

// Fly Полететь
func (f *FlyWithWings) Fly() string {
	return "I'm flying!!"
}

// FlyRocketPowered Реализация реактивного полета
type FlyRocketPowered struct{}

// Fly Полететь
func (f *FlyRocketPowered) Fly() string {
	return "I'm flying with a rocket!!"
}
