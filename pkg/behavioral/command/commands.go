package command

import "strings"

// Command Интерфейс команды.
type Command interface {
	Execute() string
	Undo() string
}

// Макрокоманда.
type macroCommand struct {
	commands []Command
}

// Execute Выполняет команду.
func (m macroCommand) Execute() string {
	log := make([]string, 0, 2)

	for i := range m.commands {
		log = append(log, m.commands[i].Execute())
	}

	return strings.Join(log, "\n")
}

// Undo Отменяет команду.
func (m macroCommand) Undo() string {
	log := make([]string, 0, 2)

	for i := range m.commands {
		log = append(log, m.commands[i].Undo())
	}

	return strings.Join(log, "\n")
}

// NewMacroCommand Создает макрокоманду.
func NewMacroCommand(commands []Command) Command {
	return macroCommand{commands}
}

// NoCommand Пустая команда.
type NoCommand struct{}

// Execute Выполняет команду.
func (c NoCommand) Execute() string {
	return ""
}

// Undo Отменяет команду.
func (c NoCommand) Undo() string {
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
	log := make([]string, 0, 2)
	log = append(log, c.garageDoor.up())
	log = append(log, c.garageDoor.lightOn())

	return strings.Join(log, "\n")
}

// Undo Отменяет команду.
func (c garageDoorUpCommand) Undo() string {
	log := make([]string, 0, 2)
	log = append(log, c.garageDoor.lightOff())
	log = append(log, c.garageDoor.down())

	return strings.Join(log, "\n")
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
	log := make([]string, 0, 2)
	log = append(log, c.garageDoor.lightOff())
	log = append(log, c.garageDoor.down())

	return strings.Join(log, "\n")
}

// Undo Отменяет команду.
func (c garageDoorDownCommand) Undo() string {
	log := make([]string, 0, 2)
	log = append(log, c.garageDoor.up())
	log = append(log, c.garageDoor.lightOn())

	return strings.Join(log, "\n")
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
	log := make([]string, 0, 3)
	log = append(log, c.stereo.on())
	log = append(log, c.stereo.setCd())
	log = append(log, c.stereo.setVolume(stereoDefaultVolume))

	return strings.Join(log, "\n")
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
	log := make([]string, 0, 3)
	log = append(log, c.stereo.on())
	log = append(log, c.stereo.setCd())
	log = append(log, c.stereo.setVolume(stereoDefaultVolume))

	return strings.Join(log, "\n")
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

// Команда "Включить телевизор".
type tvOnCommand struct {
	tv *TV
}

// Execute Выполняет команду.
func (t tvOnCommand) Execute() string {
	return t.tv.on()
}

// Undo Отменяет команду.
func (t tvOnCommand) Undo() string {
	return t.tv.off()
}

// NewTVOnCommand Создает команду "Включить телевизор".
func NewTVOnCommand(tv *TV) Command {
	return tvOnCommand{tv}
}

// Команда "Выключить телевизор".
type tvOffCommand struct {
	tv *TV
}

// Execute Выполняет команду.
func (t tvOffCommand) Execute() string {
	return t.tv.off()
}

// Undo Отменяет команду.
func (t tvOffCommand) Undo() string {
	return t.tv.on()
}

// NewTVOffCommand Создает команду "Выключить телевизор".
func NewTVOffCommand(tv *TV) Command {
	return tvOffCommand{tv}
}

// Команда "Включить джакузи".
type hotTubOnCommand struct {
	hotTub *HotTub
}

// Execute Выполняет команду.
func (t hotTubOnCommand) Execute() string {
	return t.hotTub.on()
}

// Undo Отменяет команду.
func (t hotTubOnCommand) Undo() string {
	return t.hotTub.off()
}

// NewHotTubOnCommand Создает команду "Включить джакузи".
func NewHotTubOnCommand(hotTub *HotTub) Command {
	return hotTubOnCommand{hotTub}
}

// Команда "Выключить джакузи".
type hotTubOffCommand struct {
	hotTub *HotTub
}

// Execute Выполняет команду.
func (t hotTubOffCommand) Execute() string {
	return t.hotTub.off()
}

// Undo Отменяет команду.
func (t hotTubOffCommand) Undo() string {
	return t.hotTub.on()
}

// NewHotTubOffCommand Создает команду "Выключить джакузи".
func NewHotTubOffCommand(hotTub *HotTub) Command {
	return hotTubOffCommand{hotTub}
}
