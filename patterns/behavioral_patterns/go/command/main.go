package main

import "command_method/command"

func main() {

	light := command.NewLight("Bed Room")
	thermostat := command.NewThermoStat(100)

	lightOn := command.NewLightOnCommand(light)
	lightOff := command.NewLightOffCommand(light)
	setTemp := command.NewSetTemperatureCommand(thermostat, 25)

	remote := command.NewRemoteControl()

	remote.ExecuteCommand(lightOn)
	remote.ExecuteCommand(setTemp)
	remote.ExecuteCommand(lightOff)

	remote.UndoLastCommand()
	remote.UndoLastCommand()
	remote.UndoLastCommand()
	remote.UndoLastCommand()
	remote.UndoLastCommand()

}
