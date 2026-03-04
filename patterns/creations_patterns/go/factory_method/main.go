package main

import (
	"factory_method/factory"
	"fmt"
)

func main() {
	notificationFactory := factory.NewNotificationFactory()

	notificationFactory.Register("EMAIL", func() factory.Notification {
		return factory.NewEmailNotification()
	})

	notificationFactory.Register("SMS", func() factory.Notification {
		return factory.NewSMSNotification()
	})

	notificationFactory.Register("PUSH", func() factory.Notification {
		return factory.NewPushNotification()
	})

	notifications := []string{"EMAIL", "SMS", "PUSH"}
	messages := []string{
		"Welcome to our platform!",
		"Your OTP is 123456",
		"You have a new follower!",
	}

	for i, notificationType := range notifications {
		notification, err := notificationFactory.Create(notificationType)
		if err != nil {
			fmt.Printf("Error creating notification: %v\n", err)
			continue
		}

		notification.Send(messages[i])
	}
}
