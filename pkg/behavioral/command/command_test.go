package command_test

import (
	"fmt"
	"testing"

	"github.com/Mark-Sart/go-patterns/pkg/behavioral/command"
	"github.com/stretchr/testify/assert"
)

func TestNoCommand(t *testing.T) {
	noCommand := command.NoCommand{}

	expected := [2]string{"", ""}

	testCases(t, noCommand, noCommand, expected)
}

func TestLightCommand(t *testing.T) {
	light := &command.Light{Location: "Living Room"}
	lightOn := command.NewLightOnCommand(light)
	lightOff := command.NewLightOffCommand(light)

	expected := [2]string{
		"Living Room light is on",
		"Living Room light is off",
	}

	testCases(t, lightOn, lightOff, expected)
}

func TestGarageDoorCommand(t *testing.T) {
	garageDoor := &command.GarageDoor{}
	garageDoorUp := command.NewGarageDoorUpCommand(garageDoor)
	garageDoorDown := command.NewGarageDoorDownCommand(garageDoor)

	expected := [2]string{
		"Garage door is open\nGarage light is on",
		"Garage light is off\nGarage door is close",
	}

	testCases(t, garageDoorUp, garageDoorDown, expected)
}

func TestStereoCommand(t *testing.T) {
	stereo := &command.Stereo{Location: "Living Room"}
	stereoOnWithCD := command.NewStereoOnWithCDCommand(stereo)
	stereoOff := command.NewStereoOffCommand(stereo)

	expected := [2]string{
		"Living Room stereo is on\nLiving Room stereo is set CD\nLiving Room stereo volume set to 11",
		"Living Room stereo is off",
	}

	testCases(t, stereoOnWithCD, stereoOff, expected)
}

func TestCeilingFanCommand(t *testing.T) {
	ceilingFan := &command.CeilingFan{Location: "Living Room"}
	ceilingFanHigh := command.NewCeilingFanHighCommand(ceilingFan)
	ceilingFanMedium := command.NewCeilingFanMediumCommand(ceilingFan)
	ceilingFanLow := command.NewCeilingFanLowCommand(ceilingFan)
	ceilingFanOff := command.NewCeilingFanOffCommand(ceilingFan)

	remoteControl := command.NewRemoteControl()
	remoteControl.SetCommand(command.Slot1, ceilingFanHigh, command.NoCommand{})
	remoteControl.SetCommand(command.Slot2, ceilingFanMedium, command.NoCommand{})
	remoteControl.SetCommand(command.Slot3, ceilingFanLow, command.NoCommand{})
	remoteControl.SetCommand(command.Slot4, ceilingFanOff, command.NoCommand{})

	cases := map[int]string{
		command.Slot1: "Living Room ceiling fan is on high",
		command.Slot2: "Living Room ceiling fan is on medium",
		command.Slot3: "Living Room ceiling fan is on low",
		command.Slot4: "Living Room ceiling fan is off",
	}

	for slot, expected := range cases {
		func(slot int, expected string) {
			t.Run(fmt.Sprintf("%v", slot), func(t *testing.T) {
				for innerSlot, innerExpected := range cases {
					assert.Equal(t, innerExpected, remoteControl.OnButtonWasPushed(innerSlot))
					assert.Equal(t, expected, remoteControl.OnButtonWasPushed(slot))
					assert.Equal(t, innerExpected, remoteControl.UndoButtonWasPushed())
				}
			})
		}(slot, expected)
	}

	expected := "\n------- Remote Control -------\n"
	expected += fmt.Sprintf("[slot 0] %-35T %T\n", ceilingFanHigh, command.NoCommand{})
	expected += fmt.Sprintf("[slot 1] %-35T %T\n", ceilingFanMedium, command.NoCommand{})
	expected += fmt.Sprintf("[slot 2] %-35T %T\n", ceilingFanLow, command.NoCommand{})
	expected += fmt.Sprintf("[slot 3] %-35T %T\n", ceilingFanOff, command.NoCommand{})
	expected += fmt.Sprintf("[slot 4] %-35T %T\n", command.NoCommand{}, command.NoCommand{})
	expected += fmt.Sprintf("[slot 5] %-35T %T\n", command.NoCommand{}, command.NoCommand{})
	expected += fmt.Sprintf("[slot 6] %-35T %T\n", command.NoCommand{}, command.NoCommand{})
	expected += fmt.Sprintf("[undo]   %T\n", command.NoCommand{})

	assert.Equal(t, expected, fmt.Sprint(remoteControl))
}

func TestTVCommand(t *testing.T) {
	tv := &command.TV{Location: "Living Room"}
	tvOn := command.NewTVOnCommand(tv)
	tvOff := command.NewTVOffCommand(tv)

	expected := [2]string{
		"Living Room TV is on",
		"Living Room TV is off",
	}

	testCases(t, tvOn, tvOff, expected)
}

func TestHotTub(t *testing.T) {
	hotTub := &command.HotTub{Location: "Bath Room"}
	hotTubOn := command.NewHotTubOnCommand(hotTub)
	hotTubOff := command.NewHotTubOffCommand(hotTub)

	expected := [2]string{
		"Bath Room hot tub is on",
		"Bath Room hot tub is off",
	}

	testCases(t, hotTubOn, hotTubOff, expected)
}

func TestMacroCommand(t *testing.T) {
	light := &command.Light{Location: "Living Room"}
	stereo := &command.Stereo{Location: "Living Room"}
	tv := &command.TV{Location: "Living Room"}
	hotTub := &command.HotTub{Location: "Bath Room"}

	lightOn := command.NewLightOnCommand(light)
	lightOff := command.NewLightOffCommand(light)
	stereoOn := command.NewStereoOnWithCDCommand(stereo)
	stereoOff := command.NewStereoOffCommand(stereo)
	tvOn := command.NewTVOnCommand(tv)
	tvOff := command.NewTVOffCommand(tv)
	hotTubOn := command.NewHotTubOnCommand(hotTub)
	hotTubOff := command.NewHotTubOffCommand(hotTub)

	partyOn := []command.Command{lightOn, stereoOn, tvOn, hotTubOn}
	partyOff := []command.Command{lightOff, stereoOff, tvOff, hotTubOff}

	partyOnMacro := command.NewMacroCommand(partyOn)
	partyOffMacro := command.NewMacroCommand(partyOff)

	remoteControl := command.NewRemoteControl()
	remoteControl.SetCommand(command.Slot7, partyOnMacro, partyOffMacro)

	expected := [2]string{
		"Living Room light is on\n" +
			"Living Room stereo is on\n" +
			"Living Room stereo is set CD\n" +
			"Living Room stereo volume set to 11\n" +
			"Living Room TV is on\n" +
			"Bath Room hot tub is on",
		"Living Room light is off\n" +
			"Living Room stereo is off\n" +
			"Living Room TV is off\n" +
			"Bath Room hot tub is off",
	}

	testCases(t, partyOnMacro, partyOffMacro, expected)
}

func testCases(t *testing.T, onCommand, offCommand command.Command, expected [2]string) {
	remoteControl := command.NewRemoteControl()
	remoteControl.SetCommand(command.Slot1, onCommand, offCommand)

	t.Run("Turn on", func(t *testing.T) {
		assert.Equal(t, expected[0], remoteControl.OnButtonWasPushed(command.Slot1))
		assert.Equal(t, expected[1], remoteControl.UndoButtonWasPushed())
	})
	t.Run("Turn off", func(t *testing.T) {
		assert.Equal(t, expected[1], remoteControl.OffButtonWasPushed(command.Slot1))
		assert.Equal(t, expected[0], remoteControl.UndoButtonWasPushed())
	})
}
