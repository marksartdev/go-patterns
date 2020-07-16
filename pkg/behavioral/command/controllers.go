package command

// RemoteController Интерфейс пульта дистанционного управления.
type RemoteController interface {
	SetCommand(command Command)
	ButtonWasPressed() string
}

// Простой пульт дистанционного управления.
type simpleRemoteControl struct {
	slot Command
}

// SetCommand Устанавливает команду в слот пульта.
func (c *simpleRemoteControl) SetCommand(command Command) {
	c.slot = command
}

// ButtonWasPressed Имитирует нажатие кнопки.
func (c *simpleRemoteControl) ButtonWasPressed() string {
	return c.slot.Execute()
}

// NewSimpleRemoteControl Создает простой пульт дистанционного управления.
func NewSimpleRemoteControl() RemoteController {
	return &simpleRemoteControl{}
}
