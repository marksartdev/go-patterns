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
	log := make([]string, 0, 14)
	log = append(log, "Get ready to watch a movie ...\n")
	log = append(log, h.popper.On())
	log = append(log, fmt.Sprintln(h.popper.Pop()))
	// nolint:gomnd
	log = append(log, fmt.Sprintln(h.lights.Dim(10)))
	log = append(log, fmt.Sprintln(h.screen.Down()))
	log = append(log, h.projector.On())
	log = append(log, h.projector.SetInput(h.dvd))
	log = append(log, fmt.Sprintln(h.projector.WideScreenMode()))
	log = append(log, h.amp.On())
	log = append(log, h.amp.SetDvd(h.dvd))
	log = append(log, h.amp.SetSurroundSound())
	// nolint:gomnd
	log = append(log, fmt.Sprintln(h.amp.SetVolume(5)))
	log = append(log, h.dvd.On())
	log = append(log, h.dvd.Play(movie))

	return h.returnLog(log)
}

// EndMovie Останавливает воспроизведение фильма и выключает аппаратуру.
func (h homeTheaterFacade) EndMovie() string {
	log := make([]string, 0, 9)
	log = append(log, "Shutting movie theater down ...\n")
	log = append(log, fmt.Sprintln(h.popper.Off()))
	log = append(log, fmt.Sprintln(h.lights.On()))
	log = append(log, fmt.Sprintln(h.screen.Up()))
	log = append(log, fmt.Sprintln(h.projector.Off()))
	log = append(log, fmt.Sprintln(h.amp.Off()))
	log = append(log, h.dvd.Stop())
	log = append(log, h.dvd.Eject())
	log = append(log, h.dvd.Off())

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
