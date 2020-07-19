package command

import (
	"bytes"
	"fmt"
)

// RemoteController Интерфейс пульта дистанционного управления.
type RemoteController interface {
	SetCommand(slot int, onCommand, offCommand Command)
	OnButtonWasPushed(slot int) string
	OffButtonWasPushed(slot int) string
}

// Пульт дистанционного управления.
type remoteControl struct {
	onCommands  [7]Command
	offCommands [7]Command
}

// SetCommand Устанавливает команду в слот пульта.
func (r *remoteControl) SetCommand(slot int, onCommand, offCommand Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

// OnButtonWasPushed Имитирует нажатие кнопки вкл.
func (r *remoteControl) OnButtonWasPushed(slot int) string {
	return r.onCommands[slot].Execute()
}

// OffButtonWasPushed Имитирует нажатие кнопки выкл.
func (r *remoteControl) OffButtonWasPushed(slot int) string {
	return r.offCommands[slot].Execute()
}

func (r *remoteControl) String() string {
	buffer := bytes.NewBuffer(make([]byte, 0, 10))
	buffer.WriteString("\n------- Remote Control -------\n")
	for i := 0; i < len(r.onCommands); i++ {
		buffer.WriteString(fmt.Sprintf("[slot %d] %-30T %T\n", i, r.onCommands[i], r.offCommands[i]))
	}

	return buffer.String()
}

// NewRemoteControl Создает пульт дистанционного управления.
func NewRemoteControl() RemoteController {
	controller := &remoteControl{}

	for i := 0; i < 7; i++ {
		controller.onCommands[i] = NoCommand{}
		controller.offCommands[i] = NoCommand{}
	}

	return controller
}
