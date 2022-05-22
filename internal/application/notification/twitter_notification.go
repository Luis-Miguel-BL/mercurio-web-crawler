package notification

import (
	"fmt"
	"mercurio-web-scraping/internal/config"
)

type TwitterNotification struct {
}

func NewTwitterNotification(config config.Config) *TwitterNotification {
	return &TwitterNotification{}
}

func (t *TwitterNotification) SendNotification(notification Notification) error {
	// @TODO add twitter notification
	fmt.Println("envia notificação pelo twitter")
	return nil
}
