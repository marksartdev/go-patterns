package common

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// CustomWriterSetter CustomWriter Интерфейс для установки пользовательского writer.
type CustomWriterSetter interface {
	SetWriter(writer io.Writer)
}

// CustomWriter Интерфейс для работы с пользовательским writer.
type CustomWriter interface {
	CustomWriterSetter
	Write(msg string)
}

// Обвязка для пользовательского writer.
type writer struct {
	writer io.Writer
}

// SetWriter Установить writer.
func (w *writer) SetWriter(writer io.Writer) {
	w.writer = writer
}

// Write Записать строку в writer.
func (w *writer) Write(msg string) {
	_, err := fmt.Fprintln(w.writer, msg)
	if err != nil {
		log.Error(err)
	}
}

// NewCustomWriter Создать пользовательский writer.
func NewCustomWriter() CustomWriter {
	return &writer{os.Stdout}
}
