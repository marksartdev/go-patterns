package command

import "bytes"

// Command Интерфейс команды.
type Command interface {
	Execute() string
	Undo() string
}

// Пустая команда.
type noCommand struct{}

// Execute Выполняет команду.
func (c noCommand) Execute() string {
	return ""
}

// Undo Отменяет команду.
func (c noCommand) Undo() string {
	return ""
}

// Команда "Включить свет".
type lightOnCommand struct {
	light *Light
}

// Execute Выполняет команду.
func (c lightOnCommand) Execute() string {
	return c.light.on()
}

// Undo Отменяет команду.
func (c lightOnCommand) Undo() string {
	return c.light.off()
}

// NewLightOnCommand Создает команду "Включить свет".
func NewLightOnCommand(light *Light) Command {
	return lightOnCommand{light}
}

// Команда "Выключить свет".
type lightOffCommand struct {
	light *Light
}

// Execute Выполняет команду.
func (c lightOffCommand) Execute() string {
	return c.light.off()
}

// Undo Отменяет команду.
func (c lightOffCommand) Undo() string {
	return c.light.on()
}

// NewLightOffCommand Создает команду "Выключить свет".
func NewLightOffCommand(light *Light) Command {
	return lightOffCommand{light}
}

// Команда "Открыть дверь гаража".
type garageDoorUpCommand struct {
	garageDoor *GarageDoor
}

// Execute Выполняет команду.
func (c garageDoorUpCommand) Execute() string {
	log := bytes.NewBuffer(make([]byte, 10))

	log.WriteString(c.garageDoor.up())
	log.WriteString("\n")
	log.WriteString(c.garageDoor.lightOn())

	return log.String()
}

// Undo Отменяет команду.
func (c garageDoorUpCommand) Undo() string {
	log := bytes.NewBuffer(make([]byte, 10))

	log.WriteString(c.garageDoor.lightOff())
	log.WriteString("\n")
	log.WriteString(c.garageDoor.down())

	return log.String()
}

// NewGarageDoorUpCommand Создает команду "Открыть дверь гаража".
func NewGarageDoorUpCommand(garageDoor *GarageDoor) Command {
	return garageDoorUpCommand{garageDoor}
}

// Команда "Закрыть дверь гаража".
type garageDoorDownCommand struct {
	garageDoor *GarageDoor
}

// Execute Выполняет команду.
func (c garageDoorDownCommand) Execute() string {
	log := bytes.NewBuffer(make([]byte, 10))

	log.WriteString(c.garageDoor.lightOff())
	log.WriteString("\n")
	log.WriteString(c.garageDoor.down())

	return log.String()
}

// Undo Отменяет команду.
func (c garageDoorDownCommand) Undo() string {
	log := bytes.NewBuffer(make([]byte, 10))

	log.WriteString(c.garageDoor.up())
	log.WriteString("\n")
	log.WriteString(c.garageDoor.lightOn())

	return log.String()
}

// NewGarageDoorDownCommand Создать команду "Закрыть дверь гаража".
func NewGarageDoorDownCommand(garageDoor *GarageDoor) Command {
	return garageDoorDownCommand{garageDoor}
}

// Команда "Включить стереосистему".
type stereoOnWithCDCommand struct {
	stereo *Stereo
}

// Execute Выполняет команду.
func (c stereoOnWithCDCommand) Execute() string {
	log := bytes.NewBuffer(make([]byte, 10))
	log.WriteString(c.stereo.on())
	log.WriteString("\n")
	log.WriteString(c.stereo.setCd())
	log.WriteString("\n")
	log.WriteString(c.stereo.setVolume(stereoDefaultVolume))

	return log.String()
}

// Undo Отменяет команду.
func (c stereoOnWithCDCommand) Undo() string {
	return c.stereo.off()
}

// NewStereoOnWithCDCommand Создает команду "Включить стереосистему".
func NewStereoOnWithCDCommand(stereo *Stereo) Command {
	return stereoOnWithCDCommand{stereo}
}

// Команда "Выключить стереосистему".
type stereoOffCommand struct {
	stereo *Stereo
}

// Execute Выполняет команду.
func (c stereoOffCommand) Execute() string {
	return c.stereo.off()
}

// Undo Отменяет команду.
func (c stereoOffCommand) Undo() string {
	log := bytes.NewBuffer(make([]byte, 10))
	log.WriteString(c.stereo.on())
	log.WriteString("\n")
	log.WriteString(c.stereo.setCd())
	log.WriteString("\n")
	log.WriteString(c.stereo.setVolume(stereoDefaultVolume))

	return log.String()
}

// NewStereoOffCommand Создает команду "Выключить стереосистему".
func NewStereoOffCommand(stereo *Stereo) Command {
	return stereoOffCommand{stereo}
}

// Команда "Включить высокую скорость вентилятора".
type ceilingFanHighCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

// Execute Выполняет команду.
func (c *ceilingFanHighCommand) Execute() string {
	c.prevSpeed = c.ceilingFan.getSpeed()
	return c.ceilingFan.high()
}

// Undo Отменяет команду.
func (c *ceilingFanHighCommand) Undo() string {
	switch c.prevSpeed {
	case ceilingFanHigh:
		return c.ceilingFan.high()
	case ceilingFanMedium:
		return c.ceilingFan.medium()
	case ceilingFanLow:
		return c.ceilingFan.low()
	case ceilingFanOff:
		return c.ceilingFan.off()
	default:
		return ""
	}
}

// NewCeilingFanHighCommand Создает команду "Включить высокую скорость вентилятора".
func NewCeilingFanHighCommand(ceilingFan *CeilingFan) Command {
	return &ceilingFanHighCommand{ceilingFan, ceilingFanOff}
}

// Команда "Включить среднюю скорость вентилятора".
type ceilingFanMediumCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

// Execute Выполняет команду.
func (c *ceilingFanMediumCommand) Execute() string {
	c.prevSpeed = c.ceilingFan.getSpeed()
	return c.ceilingFan.medium()
}

// Undo Отменяет команду.
func (c *ceilingFanMediumCommand) Undo() string {
	switch c.prevSpeed {
	case ceilingFanHigh:
		return c.ceilingFan.high()
	case ceilingFanMedium:
		return c.ceilingFan.medium()
	case ceilingFanLow:
		return c.ceilingFan.low()
	case ceilingFanOff:
		return c.ceilingFan.off()
	default:
		return ""
	}
}

// NewCeilingFanMediumCommand Создает команду "Включить среднюю скорость вентилятора".
func NewCeilingFanMediumCommand(ceilingFan *CeilingFan) Command {
	return &ceilingFanMediumCommand{ceilingFan, ceilingFanOff}
}

// Команда "Включить низкую скорость вентилятора".
type ceilingFanLowCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

// Execute Выполняет команду.
func (c *ceilingFanLowCommand) Execute() string {
	c.prevSpeed = c.ceilingFan.getSpeed()
	return c.ceilingFan.low()
}

// Undo Отменяет команду.
func (c *ceilingFanLowCommand) Undo() string {
	switch c.prevSpeed {
	case ceilingFanHigh:
		return c.ceilingFan.high()
	case ceilingFanMedium:
		return c.ceilingFan.medium()
	case ceilingFanLow:
		return c.ceilingFan.low()
	case ceilingFanOff:
		return c.ceilingFan.off()
	default:
		return ""
	}
}

// NewCeilingFanLowCommand Создает команду "Включить низкую скорость вентилятора".
func NewCeilingFanLowCommand(ceilingFan *CeilingFan) Command {
	return &ceilingFanLowCommand{ceilingFan, ceilingFanOff}
}

// Команда "Выключить вентилятор".
type ceilingFanOffCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

// Execute Выполняет команду.
func (c *ceilingFanOffCommand) Execute() string {
	c.prevSpeed = c.ceilingFan.getSpeed()
	return c.ceilingFan.off()
}

// Undo Отменяет команду.
func (c *ceilingFanOffCommand) Undo() string {
	switch c.prevSpeed {
	case ceilingFanHigh:
		return c.ceilingFan.high()
	case ceilingFanMedium:
		return c.ceilingFan.medium()
	case ceilingFanLow:
		return c.ceilingFan.low()
	case ceilingFanOff:
		return c.ceilingFan.off()
	default:
		return ""
	}
}

// NewCeilingFanOffCommand Создает команду "Выключить вентилятор".
func NewCeilingFanOffCommand(ceilingFan *CeilingFan) Command {
	return &ceilingFanOffCommand{ceilingFan, ceilingFanOff}
}
