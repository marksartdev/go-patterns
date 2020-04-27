package strategy

import (
	"strings"
	"unicode"
)

// Quacker Интерфейс способности крякать
type Quacker interface {
	quack(count int) string
}

// MuteQuack Реализация для молчаливых уток
type MuteQuack struct{}

// Крякнуть.
func (m *MuteQuack) quack(_ int) string {
	return "<< Silence >>"
}

// Quack Реализация для крякающих уток
type Quack struct{}

// Крякнуть.
func (q *Quack) quack(count int) string {
	speech := make([]string, count)
	for i := 0; i < count; i++ {
		speech[i] = "quack"
	}

	runes := []rune(strings.Join(speech, "-"))
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}

// Squeak Реализация для пищащих уток
type Squeak struct{}

// Крякнуть.
func (s *Squeak) quack(count int) string {
	speech := make([]string, count)
	for i := 0; i < count; i++ {
		speech[i] = "squeak"
	}

	runes := []rune(strings.Join(speech, "-"))
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
