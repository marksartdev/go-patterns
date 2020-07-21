package command

import (
	"bytes"
	"fmt"
)

const (
	// Slot1 Ячейка 1.
	Slot1 = iota
	// Slot2 Ячейка 2.
	Slot2
	// Slot3 Ячейка 3.
	Slot3
	// Slot4 Ячейка 4.
	Slot4
	// Slot5 Ячейка 5.
	Slot5
	// Slot6 Ячейка 6.
	Slot6
	// Slot7 Ячейка 7.
	Slot7
)

// RemoteController Интерфейс пульта дистанционного управления.
type RemoteController interface {
	SetCommand(slot int, onCommand, offCommand Command)
	OnButtonWasPushed(slot int) string
	OffButtonWasPushed(slot int) string
	UndoButtonWasPushed() string
}

// Пульт дистанционного управления.
type remoteControl struct {
	onCommands  [7]Command
	offCommands [7]Command
	undoCommand Command
}

// SetCommand Устанавливает команду в слот пульта.
func (r *remoteControl) SetCommand(slot int, onCommand, offCommand Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

// OnButtonWasPushed Имитирует нажатие кнопки вкл.
func (r *remoteControl) OnButtonWasPushed(slot int) string {
	result := r.onCommands[slot].Execute()
	r.undoCommand = r.onCommands[slot]

	return result
}

// OffButtonWasPushed Имитирует нажатие кнопки выкл.
func (r *remoteControl) OffButtonWasPushed(slot int) string {
	result := r.offCommands[slot].Execute()
	r.undoCommand = r.offCommands[slot]

	return result
}

// UndoButtonWasPushed Имитирует нажатие кнопки отмены.
func (r *remoteControl) UndoButtonWasPushed() string {
	result := r.undoCommand.Undo()
	r.undoCommand = noCommand{}

	return result
}

func (r *remoteControl) String() string {
	buffer := bytes.NewBuffer(make([]byte, 0, 10))
	buffer.WriteString("\n------- Remote Control -------\n")

	for i := 0; i < len(r.onCommands); i++ {
		buffer.WriteString(fmt.Sprintf("[slot %d] %-30T %T\n", i, r.onCommands[i], r.offCommands[i]))
	}

	buffer.WriteString(fmt.Sprintf("[undo]   %T\n", r.undoCommand))

	return buffer.String()
}

// NewRemoteControl Создает пульт дистанционного управления.
func NewRemoteControl() RemoteController {
	controller := &remoteControl{}

	for i := 0; i < 7; i++ {
		controller.onCommands[i] = noCommand{}
		controller.offCommands[i] = noCommand{}
	}

	controller.undoCommand = noCommand{}

	return controller
}
