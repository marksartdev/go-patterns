// Package facade Паттерн "Фасад".
package facade

import (
	"fmt"
	"strings"
)

// HomeTheaterFacade Интерфейс фасада домашнего кинотеатра.
type HomeTheaterFacade interface {
	WatchMovie(movie string) string
	EndMovie() string
}

// Фасад домашнего кинотеатра.
type homeTheaterFacade struct {
	lights    TheaterLights
	screen    Screen
	popper    PopcornPopper
	amp       *Amplifier
	dvd       *DvdPlayer
	projector *Projector
}

// WatchMovie Подготавливает домашний кинотеатр к просмотру фильма.
func (h homeTheaterFacade) WatchMovie(movie string) string {
	log := []string{
		"Get ready to watch a movie ...\n",
		h.popper.On(),
		fmt.Sprintln(h.popper.Pop()),
		// nolint:gomnd // Example
		fmt.Sprintln(h.lights.Dim(10)),
		fmt.Sprintln(h.screen.Down()),
		h.projector.On(),
		h.projector.SetInput(h.dvd),
		fmt.Sprintln(h.projector.WideScreenMode()),
		h.amp.On(),
		h.amp.SetDvd(h.dvd),
		h.amp.SetSurroundSound(),
		// nolint:gomnd // Example
		fmt.Sprintln(h.amp.SetVolume(5)),
		h.dvd.On(),
		h.dvd.Play(movie),
	}

	return h.returnLog(log)
}

// EndMovie Останавливает воспроизведение фильма и выключает аппаратуру.
func (h homeTheaterFacade) EndMovie() string {
	log := []string{
		"Shutting movie theater down ...\n",
		fmt.Sprintln(h.popper.Off()),
		fmt.Sprintln(h.lights.On()),
		fmt.Sprintln(h.screen.Up()),
		fmt.Sprintln(h.projector.Off()),
		fmt.Sprintln(h.amp.Off()),
		h.dvd.Stop(),
		h.dvd.Eject(),
		h.dvd.Off(),
	}

	return h.returnLog(log)
}

func (h homeTheaterFacade) returnLog(log []string) string {
	return strings.Join(log, "\n")
}

// NewHomeTheater Создает фасад домашнего кинотеатра.
func NewHomeTheater() HomeTheaterFacade {
	return homeTheaterFacade{
		lights:    TheaterLights{},
		screen:    Screen{},
		popper:    PopcornPopper{},
		amp:       &Amplifier{},
		dvd:       &DvdPlayer{},
		projector: &Projector{},
	}
}
