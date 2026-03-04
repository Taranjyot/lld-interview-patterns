package factory

import "fmt"

type PushNotification struct{}

func NewPushNotification() *PushNotification {
	return &PushNotification{}
}

func (this *PushNotification) Send(message string) {
	fmt.Printf("Sending SMS: %s \n", message)
}
