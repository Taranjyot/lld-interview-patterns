package factory

import "fmt"

type SimpleNotificationFactory struct{}

func NewSimpleNotificationFactory() *SimpleNotificationFactory {
	return &SimpleNotificationFactory{}
}

func (f *SimpleNotificationFactory) CreateNotification(notificationType string) (Notification, error) {
	switch notificationType {
	case "EMAIL":
		return NewEmailNotification(), nil
	case "SMS":
		return NewSMSNotification(), nil
	case "PUSH":
		return NewPushNotification(), nil
	default:
		return nil, fmt.Errorf("unknown notification type: %s", notificationType)
	}
}
