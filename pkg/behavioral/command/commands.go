package command

// Command Интерфейс команды.
type Command interface {
	Execute() string
}

// Команда "Включить свет".
type lightOnCommand struct {
	light Light
}

// Execute Выполняет команду.
func (c lightOnCommand) Execute() string {
	return c.light.on()
}

// NewLightOnCommand Создает команду "Включить свет".
func NewLightOnCommand(light Light) Command {
	return lightOnCommand{light}
}

// Команда "Открыть дверь гаража".
type garageDoorOpenCommand struct {
	garageDoor GarageDoor
}

// Execute Выполняет команду.
func (c garageDoorOpenCommand) Execute() string {
	log := c.garageDoor.up()
	log += "\n"
	log += c.garageDoor.lightOn()

	return log
}

// NewGarageDoorOpenCommand Создает команду "Открыть дверь гаража".
func NewGarageDoorOpenCommand(garageDoor GarageDoor) Command {
	return garageDoorOpenCommand{garageDoor}
}
