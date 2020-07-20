package command

import "bytes"

// Command Интерфейс команды.
type Command interface {
	Execute() string
}

// NoCommand Пустая команда.
type NoCommand struct {
}

// Execute Выполняет команду.
func (c NoCommand) Execute() string {
	return ""
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

// Команда "Выключить свет".
type lightOffCommand struct {
	light Light
}

// Execute Выполняет команду.
func (c lightOffCommand) Execute() string {
	return c.light.off()
}

// NewLightOffCommand Создает команду "Выключить свет".
func NewLightOffCommand(light Light) Command {
	return lightOffCommand{light}
}

// Команда "Открыть дверь гаража".
type garageDoorUpCommand struct {
	garageDoor GarageDoor
}

// Execute Выполняет команду.
func (c garageDoorUpCommand) Execute() string {
	log := bytes.NewBuffer(make([]byte, 10))

	log.WriteString(c.garageDoor.up())
	log.WriteString("\n")
	log.WriteString(c.garageDoor.lightOn())

	return log.String()
}

// NewGarageDoorUpCommand Создает команду "Открыть дверь гаража".
func NewGarageDoorUpCommand(garageDoor GarageDoor) Command {
	return garageDoorUpCommand{garageDoor}
}

// Команда "Закрыть дверь гаража".
type garageDoorDownCommand struct {
	garageDoor GarageDoor
}

// Execute Выполняет команду.
func (c garageDoorDownCommand) Execute() string {
	log := bytes.NewBuffer(make([]byte, 10))

	log.WriteString(c.garageDoor.lightOff())
	log.WriteString("\n")
	log.WriteString(c.garageDoor.down())

	return log.String()
}

// NewGarageDoorDownCommand Создать команду "Закрыть дверь гаража".
func NewGarageDoorDownCommand(garageDoor GarageDoor) Command {
	return garageDoorDownCommand{garageDoor}
}

// Команда "Включить стереосистему".
type stereoOnWithCDCommand struct {
	stereo Stereo
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

// NewStereoOnWithCDCommand Создает команду "Включить стереосистему".
func NewStereoOnWithCDCommand(stereo Stereo) Command {
	return stereoOnWithCDCommand{stereo}
}

// Команда "Выключить стереосистему".
type stereoOffCommand struct {
	stereo Stereo
}

// Execute Выполняет команду.
func (c stereoOffCommand) Execute() string {
	return c.stereo.off()
}

// NewStereoOffCommand Создает команду "Выключить стереосистему".
func NewStereoOffCommand(stereo Stereo) Command {
	return stereoOffCommand{stereo}
}

// Команда "Включить вентилятор".
type ceilingFanOnCommand struct {
	ceilingFan CeilingFan
}

// Execute Выполняет команду.
func (c ceilingFanOnCommand) Execute() string {
	return c.ceilingFan.high()
}

// NewCeilingFanOnCommand Создает команду "Включить вентилятор".
func NewCeilingFanOnCommand(ceilingFan CeilingFan) Command {
	return ceilingFanOnCommand{ceilingFan}
}

// Команда "Выключить вентилятор".
type ceilingFanOffCommand struct {
	ceilingFan CeilingFan
}

// Execute Выполняет команду.
func (c ceilingFanOffCommand) Execute() string {
	return c.ceilingFan.off()
}

// NewCeilingFanOffCommand Создает команду "Выключить вентилятор".
func NewCeilingFanOffCommand(ceilingFan CeilingFan) Command {
	return ceilingFanOffCommand{ceilingFan}
}
