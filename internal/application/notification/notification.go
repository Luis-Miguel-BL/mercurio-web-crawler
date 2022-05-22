package notification

type NotificationHandler interface {
	SendNotification(notification Notification) error
}

type NotificationChannel string

var NotificationChannelTwitter NotificationChannel = "twitter"

type Notification struct {
	Channels    []NotificationChannel
	Destination string
	Message     string
}
