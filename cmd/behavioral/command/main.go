package main

import (
	"fmt"

	"github.com/Mark-Sart/go-patterns/pkg/behavioral/command"
)

func main() {
	var (
		livingRoomLight, kitchenLight *command.Light
		ceilingFan                    *command.CeilingFan
		garageDoor                    *command.GarageDoor
		stereo                        *command.Stereo
	)

	livingRoomLight = &command.Light{Location: "Living Room"}
	kitchenLight = &command.Light{Location: "Kitchen"}
	ceilingFan = &command.CeilingFan{Location: "Living Room"}
	garageDoor = &command.GarageDoor{}
	stereo = &command.Stereo{Location: "Living Room"}

	test0(livingRoomLight, kitchenLight, ceilingFan, garageDoor, stereo)
	test1(livingRoomLight)
	test2(ceilingFan)
}

func test0(
	livingRoomLight, kitchenLight *command.Light,
	ceilingFan *command.CeilingFan,
	garageDoor *command.GarageDoor,
	stereo *command.Stereo,
) {
	var (
		remoteControl                         command.RemoteController
		livingRoomLightOn, livingRoomLightOff command.Command
		kitchenLightOn, kitchenLightOff       command.Command
		ceilingFanHigh, ceilingFanOff         command.Command
		garageDoorUp, garageDoorDown          command.Command
		stereoOnWithCD, stereoOff             command.Command
	)

	livingRoomLightOn = command.NewLightOnCommand(livingRoomLight)
	livingRoomLightOff = command.NewLightOffCommand(livingRoomLight)
	kitchenLightOn = command.NewLightOnCommand(kitchenLight)
	kitchenLightOff = command.NewLightOffCommand(kitchenLight)
	ceilingFanHigh = command.NewCeilingFanHighCommand(ceilingFan)
	ceilingFanOff = command.NewCeilingFanOffCommand(ceilingFan)
	garageDoorUp = command.NewGarageDoorUpCommand(garageDoor)
	garageDoorDown = command.NewGarageDoorDownCommand(garageDoor)
	stereoOnWithCD = command.NewStereoOnWithCDCommand(stereo)
	stereoOff = command.NewStereoOffCommand(stereo)

	remoteControl = command.NewRemoteControl()
	remoteControl.SetCommand(command.Slot1, livingRoomLightOn, livingRoomLightOff)
	remoteControl.SetCommand(command.Slot2, kitchenLightOn, kitchenLightOff)
	remoteControl.SetCommand(command.Slot3, ceilingFanHigh, ceilingFanOff)
	remoteControl.SetCommand(command.Slot4, garageDoorUp, garageDoorDown)
	remoteControl.SetCommand(command.Slot5, stereoOnWithCD, stereoOff)

	fmt.Println(remoteControl)
	fmt.Println("----------- Test 0 -----------")

	for i := 0; i < 7; i++ {
		resultOn := remoteControl.OnButtonWasPushed(i)
		if resultOn != "" {
			fmt.Println(resultOn)
		}

		resultOff := remoteControl.OffButtonWasPushed(i)
		if resultOff != "" {
			fmt.Println(resultOff)
		}

		if len(resultOn)+len(resultOff) > 0 {
			fmt.Println()
		}
	}
}

func test1(livingRoomLight *command.Light) {
	var (
		remoteControl                         command.RemoteController
		livingRoomLightOn, livingRoomLightOff command.Command
	)

	livingRoomLightOn = command.NewLightOnCommand(livingRoomLight)
	livingRoomLightOff = command.NewLightOffCommand(livingRoomLight)

	remoteControl = command.NewRemoteControl()
	remoteControl.SetCommand(command.Slot1, livingRoomLightOn, livingRoomLightOff)

	fmt.Println("----------- Test 1 -----------")
	fmt.Println(remoteControl.OnButtonWasPushed(command.Slot1))
	fmt.Println(remoteControl.OffButtonWasPushed(command.Slot1))
	fmt.Println(remoteControl)
	fmt.Println(remoteControl.UndoButtonWasPushed())
	fmt.Println()
	fmt.Println(remoteControl.OffButtonWasPushed(command.Slot1))
	fmt.Println(remoteControl.OnButtonWasPushed(command.Slot1))
	fmt.Println(remoteControl)
	fmt.Println(remoteControl.UndoButtonWasPushed())
}

func test2(ceilingFan *command.CeilingFan) {
	var (
		remoteControl                   command.RemoteController
		ceilingFanLow, ceilingFanMedium command.Command
		ceilingFanHigh, ceilingFanOff   command.Command
	)

	ceilingFanLow = command.NewCeilingFanLowCommand(ceilingFan)
	ceilingFanMedium = command.NewCeilingFanMediumCommand(ceilingFan)
	ceilingFanHigh = command.NewCeilingFanHighCommand(ceilingFan)
	ceilingFanOff = command.NewCeilingFanOffCommand(ceilingFan)

	remoteControl = command.NewRemoteControl()
	remoteControl.SetCommand(command.Slot1, ceilingFanLow, ceilingFanOff)
	remoteControl.SetCommand(command.Slot2, ceilingFanMedium, ceilingFanOff)
	remoteControl.SetCommand(command.Slot3, ceilingFanHigh, ceilingFanOff)

	fmt.Println("----------- Test 2 -----------")
	fmt.Println(remoteControl.OnButtonWasPushed(command.Slot2))
	fmt.Println(remoteControl.OffButtonWasPushed(command.Slot2))
	fmt.Println(remoteControl)
	fmt.Println(remoteControl.UndoButtonWasPushed())
	fmt.Println(remoteControl.OnButtonWasPushed(command.Slot3))
	fmt.Println(remoteControl)
	fmt.Println(remoteControl.UndoButtonWasPushed())
}
