package facade_test

import (
	"fmt"
	"testing"

	"github.com/Mark-Sart/go-patterns/pkg/structural/facade"
	"github.com/stretchr/testify/assert"
)

func TestHomeTheaterFacade(t *testing.T) {
	homeTheater := facade.NewHomeTheater()

	lights := facade.TheaterLights{}
	screen := facade.Screen{}
	popper := facade.PopcornPopper{}
	amp := &facade.Amplifier{}
	dvd := &facade.DvdPlayer{}
	projector := &facade.Projector{}

	t.Run("WatchMovie", func(t *testing.T) {
		expected := "Get ready to watch a movie ...\n\n"
		expected += fmt.Sprintln(popper.On())
		expected += fmt.Sprintf("%s\n\n", popper.Pop())
		expected += fmt.Sprintf("%s\n\n", lights.Dim(10))
		expected += fmt.Sprintf("%s\n\n", screen.Down())
		expected += fmt.Sprintln(projector.On())
		expected += fmt.Sprintln(projector.SetInput(dvd))
		expected += fmt.Sprintf("%s\n\n", projector.WideScreenMode())
		expected += fmt.Sprintln(amp.On())
		expected += fmt.Sprintln(amp.SetDvd(dvd))
		expected += fmt.Sprintln(amp.SetSurroundSound())
		expected += fmt.Sprintf("%s\n\n", amp.SetVolume(5))
		expected += fmt.Sprintln(dvd.On())
		expected += dvd.Play("The Big Bang Theory")

		actual := homeTheater.WatchMovie("The Big Bang Theory")
		assert.Equal(t, expected, actual)
	})

	t.Run("EndMovie", func(t *testing.T) {
		expected := "Shutting movie theater down ...\n\n"
		expected += fmt.Sprintf("%s\n\n", popper.Off())
		expected += fmt.Sprintf("%s\n\n", lights.On())
		expected += fmt.Sprintf("%s\n\n", screen.Up())
		expected += fmt.Sprintf("%s\n\n", projector.Off())
		expected += fmt.Sprintf("%s\n\n", amp.Off())
		expected += fmt.Sprintln(dvd.Stop())
		expected += fmt.Sprintln(dvd.Eject())
		expected += dvd.Off()

		actual := homeTheater.EndMovie()
		assert.Equal(t, expected, actual)
	})
}
