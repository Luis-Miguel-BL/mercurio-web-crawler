package entities

import (
	"errors"
	"fmt"
	"mercurio-web-scraping/internal/application/notification"
)

var NotificationCollectionName = "notification"

type Notification struct {
	Base          `bson:",inline"`
	Channel       notification.NotificationChannel `json:"channel" bson:"channel"`
	Contact       string                           `json:"contact" bson:"contact"`
	HarvestTarget HarvestType                      `json:"harvest_target" bson:"harvest_target"`
}

func (n *Notification) Validate() (err error) {
	if n.Channel == "" {
		return errors.New("invalid channel")
	}
	if n.HarvestTarget == "" {
		return errors.New("invalid harvest target")
	}
	if n.Contact == "" {
		return errors.New("invalid contact")
	}

	return nil

}
func (n *Notification) BuildNotification(harvest Harvest) notification.Notification {
	message := fmt.Sprintf(" Novo(a) %s encontrada no site %s", translateHarvestType[harvest.HarvestType], harvest.PageLink)

	return notification.Notification{
		Channels:    []notification.NotificationChannel{n.Channel},
		Destination: n.Contact,
		Message:     message,
	}

}

var translateHarvestType = map[HarvestType]string{
	HarvestBuilding: "moradia",
}
