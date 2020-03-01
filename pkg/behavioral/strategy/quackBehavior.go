package strategy

import (
	"strings"
	"unicode"
)

// QuackBehavior Интерфейс способности крякать
type QuackBehavior interface {
	Quack(count int) string
}

// MuteQuack Реализация для молчаливых уток
type MuteQuack struct{}

// Quack Крякнуть
func (q *MuteQuack) Quack(_ int) string {
	return "<< Silence >>"
}

// Quack Реализация для крякающих уток
type Quack struct{}

// Quack Крякнуть
func (q *Quack) Quack(count int) string {
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

// Quack Крякнуть
func (q *Squeak) Quack(count int) string {
	speech := make([]string, count)
	for i := 0; i < count; i++ {
		speech[i] = "squeak"
	}

	runes := []rune(strings.Join(speech, "-"))
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
