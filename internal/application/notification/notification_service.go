package notification

import "mercurio-web-scraping/internal/config"

type NotificationService struct {
	NotificationHandlers map[NotificationChannel]NotificationHandler
}

func NewNotificationService(config config.Config) *NotificationService {
	return &NotificationService{NotificationHandlers: map[NotificationChannel]NotificationHandler{
		NotificationChannelTwitter:  NewTwitterNotification(config),
		NotificationChannelWhatsApp: NewWhatsappNotification(config),
	}}
}

func (n *NotificationService) Notificate(notification Notification) (err error) {
	for _, channel := range notification.Channels {
		err = n.NotificationHandlers[channel].SendNotification(notification)
		if err != nil {
			return err
		}
	}
	return nil
}
