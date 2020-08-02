package main

import (
	"fmt"

	"github.com/Mark-Sart/go-patterns/pkg/structural/facade"
)

// nolint:gomnd
func main() {
	popper := facade.PopcornPopper{}
	lights := facade.TheaterLights{}
	screen := facade.Screen{}
	projection := &facade.Projector{}
	amp := &facade.Amplifier{}
	dvd := &facade.DvdPlayer{}

	fmt.Println()
	fmt.Println(popper.On())
	fmt.Println(popper.Pop())
	fmt.Println()

	fmt.Println(lights.On())
	fmt.Println(lights.Dim(10))
	fmt.Println()

	fmt.Println(screen.Down())
	fmt.Println()

	fmt.Println(projection.On())
	fmt.Println(projection.SetInput(dvd))
	fmt.Println(projection.WideScreenMode())
	fmt.Println()

	fmt.Println(amp.On())
	fmt.Println(amp.SetDvd(dvd))
	fmt.Println(amp.SetSurroundSound())
	fmt.Println(amp.SetVolume(5))
	fmt.Println()

	fmt.Println(dvd.On())
	fmt.Println(dvd.Play("The Big Bang Theory"))
	fmt.Println()
}
