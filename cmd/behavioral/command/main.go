package main

import (
	"fmt"

	"github.com/Mark-Sart/go-patterns/pkg/behavioral/command"
)

func main() {
	var (
		remote              command.RemoteController
		light               command.Light
		garageDoor          command.GarageDoor
		lightOn, garageOpen command.Command
	)

	remote = command.NewSimpleRemoteControl()
	light = command.Light{}
	garageDoor = command.GarageDoor{}
	lightOn = command.NewLightOnCommand(light)
	garageOpen = command.NewGarageDoorOpenCommand(garageDoor)

	remote.SetCommand(lightOn)
	fmt.Println(remote.ButtonWasPressed())
	remote.SetCommand(garageOpen)
	fmt.Println(remote.ButtonWasPressed())
}
