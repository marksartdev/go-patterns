package facade

import "fmt"

// PopcornPopper Аппарат для попкорна.
type PopcornPopper struct{}

// On Включает аппарат для попкорна.
func (PopcornPopper) On() string {
	return "Popcorn Popper on"
}

// Off Выключает аппарат для попкорна.
func (PopcornPopper) Off() string {
	return "Popcorn Popper off"
}

// Pop Запускает приготовление попкорна.
func (PopcornPopper) Pop() string {
	return "Popcorn Popper popping popcorn!"
}

// TheaterLights Освещение в кинотеатре.
type TheaterLights struct{}

// On Включает освещение.
func (TheaterLights) On() string {
	return "Theater Ceiling Lights on"
}

// Off Выключает освещение.
func (TheaterLights) Off() string {
	return "Theater Ceiling Lights off"
}

// Dim Устанавливает уровень освещения.
// nolint:gomnd
func (TheaterLights) Dim(level int) string {
	if level > 100 {
		level = 100
	}

	if level < 0 {
		level = 0
	}

	return fmt.Sprintf("Theater Ceiling Lights dimming to %d%%", level)
}

// Screen Экран.
type Screen struct{}

// Up Поднимает экран.
func (Screen) Up() string {
	return "Theater Screen going up"
}

// Down Опускает экран.
func (Screen) Down() string {
	return "Theater Screen going down"
}

// DvdPlayer DVD-плейер.
type DvdPlayer struct {
	movie string
}

// On Включает DVD-плейер.
func (d *DvdPlayer) On() string {
	return "DVD Player on"
}

// Off Выключает DVD-плейер.
func (d *DvdPlayer) Off() string {
	return "DVD Player off"
}

// Eject Извлекает диск.
func (d *DvdPlayer) Eject() string {
	return "DVD Player eject"
}

// Play Воспроизводит фильм.
func (d *DvdPlayer) Play(movie string) string {
	d.movie = movie
	return fmt.Sprintf("DVD Player playing %q", d.movie)
}

// Stop Останавливает воспроизведение.
func (d *DvdPlayer) Stop() string {
	return fmt.Sprintf("DVD Player stopped %q", d.movie)
}

// Projector Проектор.
type Projector struct {
	dvdPlayer *DvdPlayer
}

// On Вкючает проектор.
func (p *Projector) On() string {
	return "Projector on"
}

// Off Выключает проектор.
func (p *Projector) Off() string {
	return "Projector off"
}

// SetInput Устанавливает источник сигнала.
func (p *Projector) SetInput(dvdPlayer *DvdPlayer) string {
	p.dvdPlayer = dvdPlayer
	return "Projector setting DVD Player to input"
}

// WideScreenMode Включает широкоформатный режим.
func (p *Projector) WideScreenMode() string {
	return "Projector in widescreen mode (16x9 aspect ratio)"
}

// Amplifier Усилитель.
type Amplifier struct {
	dvdPlayer *DvdPlayer
}

// On Включает усилитель.
func (a *Amplifier) On() string {
	return "Amplifier on"
}

// Off Выключает усилитель.
func (a *Amplifier) Off() string {
	return "Amplifier off"
}

// SetDvd Устанавливает источник сигнала.
func (a *Amplifier) SetDvd(dvdPlayer *DvdPlayer) string {
	a.dvdPlayer = dvdPlayer
	return "Amplifier setting DVD Player to input"
}

// SetSurroundSound Включает режим окружающего звука.
func (a *Amplifier) SetSurroundSound() string {
	return "Amplifier surround sound on (5 speakers, 1 subwoofer)"
}

// SetVolume Устанавливает громкость.
func (a *Amplifier) SetVolume(level int) string {
	return fmt.Sprintf("Amplifier setting volume to %d", level)
}
