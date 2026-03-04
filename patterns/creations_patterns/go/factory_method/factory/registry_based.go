package factory

import "fmt"

type NotificationFactory struct {
	creators map[string]func() Notification
}

func NewNotificationFactory() *NotificationFactory {
	return &NotificationFactory{
		creators: make(map[string]func() Notification),
	}
}

func (f *NotificationFactory) Register(kind string, creator func() Notification) {
	f.creators[kind] = creator
}

func (f *NotificationFactory) Create(kind string) (Notification, error) {
	creator, ok := f.creators[kind]

	if !ok {
		return nil, fmt.Errorf("unknown notification type: %s", kind)
	}

	return creator(), nil
}
