package command

import "fmt"

type RemoteControl struct {
	history []Command
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		history: []Command{},
	}
}

func (this *RemoteControl) ExecuteCommand(command Command) {
	command.execute()
	this.history = append(this.history, command)
}

func (this *RemoteControl) UndoLastCommand() {
	if len(this.history) > 0 {
		c := this.history[len(this.history)-1]
		c.undo()
		this.history = this.history[:len(this.history)-1]
	} else {
		fmt.Println("No history found!!")
	}
}
