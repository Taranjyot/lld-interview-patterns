package command

import "fmt"

type Thermostat struct {
	currentTemprature int
}

func NewThermoStat(temp int) *Thermostat {
	return &Thermostat{
		currentTemprature: temp,
	}
}

func (this *Thermostat) setTemprature(temp int) {
	this.currentTemprature = temp
	fmt.Printf("Thermostat set to %d \n", this.currentTemprature)
}

func (this *Thermostat) getTemprature() int {
	fmt.Printf("Thermostat value is %d \n", this.currentTemprature)
	return this.currentTemprature
}
