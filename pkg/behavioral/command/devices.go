package command

// Light Светильник.
type Light struct{}

// Включает свет.
func (Light) on() string {
	return "Light is On"
}

//// Выключает свет.
//func (Light) off() string {
//	return "Light is Off"
//}

// GarageDoor Дверь гаража.
type GarageDoor struct{}

// Открывает дверь гаража.
func (GarageDoor) up() string {
	return "Garage Door is Open"
}

//// Закрывает дверь гаража.
//func (GarageDoor) down() string {
//	return "Garage Door is Close"
//}

// Включает свет в гараже.
func (GarageDoor) lightOn() string {
	return "Garage Light is On"
}

//// Выключает свет в гараже.
//func (GarageDoor) lightOff() string {
//	return "Garage Light is Off"
//}
