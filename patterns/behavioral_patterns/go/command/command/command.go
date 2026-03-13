package command

type Command interface {
	execute()
	undo()
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{
		light: light,
	}
}

func (this *LightOnCommand) execute() {
	this.light.On()
}

func (this *LightOnCommand) undo() {
	this.light.Off()
}

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{
		light: light,
	}
}

func (this *LightOffCommand) execute() {
	this.light.Off()
}

func (this *LightOffCommand) undo() {
	this.light.On()
}

type SetTemperatureCommand struct {
	thermostat          *Thermostat
	newTemperature      int
	previousTemperature int
}

func NewSetTemperatureCommand(thermostat *Thermostat, newTemperature int) *SetTemperatureCommand {
	return &SetTemperatureCommand{
		thermostat:     thermostat,
		newTemperature: newTemperature,
	}
}

func (this *SetTemperatureCommand) execute() {
	this.previousTemperature = this.thermostat.getTemprature()
	this.thermostat.setTemprature(this.newTemperature)
}

func (this *SetTemperatureCommand) undo() {
	this.thermostat.setTemprature(this.previousTemperature)
}
