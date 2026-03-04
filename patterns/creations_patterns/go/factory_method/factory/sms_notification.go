package factory

import "fmt"

type SMSNotificaiton struct{}

func NewSMSNotification() *SMSNotificaiton {
	return &SMSNotificaiton{}
}

func (this *SMSNotificaiton) Send(message string) {
	fmt.Printf("Sending SMS: %s \n", message)
}
