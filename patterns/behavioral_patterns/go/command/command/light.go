package command

import "fmt"

type Light struct {
	location string
	isOn     bool
}

func NewLight(location string) *Light {
	return &Light{
		location: location,
	}
}

func (this *Light) On() {
	this.isOn = true
	fmt.Printf("%s light is On \n", this.location)
}

func (this *Light) Off() {
	this.isOn = false
	fmt.Printf("%s light is OFF \n", this.location)
}
