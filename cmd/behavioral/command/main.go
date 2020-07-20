package main

import (
	"fmt"

	"github.com/Mark-Sart/go-patterns/pkg/behavioral/command"
)

func main() {
	var (
		remoteControl                         command.RemoteController
		livingRoomLight, kitchenLight         command.Light
		ceilingFan                            command.CeilingFan
		garageDoor                            command.GarageDoor
		stereo                                command.Stereo
		livingRoomLightOn, livingRoomLightOff command.Command
		kitchenLightOn, kitchenLightOff       command.Command
		ceilingFanOn, ceilingFanOff           command.Command
		garageDoorUp, garageDoorDown          command.Command
		stereoOnWithCD, stereoOff             command.Command
	)

	remoteControl = command.NewRemoteControl()
	livingRoomLight = command.Light{Location: "Living Room"}
	kitchenLight = command.Light{Location: "Kitchen"}
	ceilingFan = command.CeilingFan{Location: "Living Room"}
	garageDoor = command.GarageDoor{}
	stereo = command.Stereo{Location: "Living Room"}

	livingRoomLightOn = command.NewLightOnCommand(livingRoomLight)
	livingRoomLightOff = command.NewLightOffCommand(livingRoomLight)
	kitchenLightOn = command.NewLightOnCommand(kitchenLight)
	kitchenLightOff = command.NewLightOffCommand(kitchenLight)
	ceilingFanOn = command.NewCeilingFanOnCommand(ceilingFan)
	ceilingFanOff = command.NewCeilingFanOffCommand(ceilingFan)
	garageDoorUp = command.NewGarageDoorUpCommand(garageDoor)
	garageDoorDown = command.NewGarageDoorDownCommand(garageDoor)
	stereoOnWithCD = command.NewStereoOnWithCDCommand(stereo)
	stereoOff = command.NewStereoOffCommand(stereo)

	remoteControl.SetCommand(command.Slot1, livingRoomLightOn, livingRoomLightOff)
	remoteControl.SetCommand(command.Slot2, kitchenLightOn, kitchenLightOff)
	remoteControl.SetCommand(command.Slot3, ceilingFanOn, ceilingFanOff)
	remoteControl.SetCommand(command.Slot4, garageDoorUp, garageDoorDown)
	remoteControl.SetCommand(command.Slot5, stereoOnWithCD, stereoOff)

	fmt.Println(remoteControl)

	for i := 0; i < 7; i++ {
		result := remoteControl.OnButtonWasPushed(i)
		if result != "" {
			fmt.Println(result)
		}

		result = remoteControl.OffButtonWasPushed(i)
		if result != "" {
			fmt.Println(result)
		}
	}
}
