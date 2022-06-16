package notification

import (
	"fmt"
	"mercurio-web-scraping/internal/config"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterNotification struct {
	client *twitter.Client
}

func NewTwitterNotification(config config.Config) *TwitterNotification {
	cfg := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecretKey)
	token := oauth1.NewToken(config.TokenKey, config.TokenScretKey)
	httpClient := cfg.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	return &TwitterNotification{client: client}
}

func (t *TwitterNotification) SendNotification(notification Notification) error {
	fmt.Println("enviou notificação")
	a, b, c := t.client.DirectMessages.EventsNew(&twitter.DirectMessageEventsNewParams{
		Event: &twitter.DirectMessageEvent{
			Message: &twitter.DirectMessageEventMessage{
				SenderID: "24345266",
				Target:   &twitter.DirectMessageTarget{RecipientID: "758783898045659279"},
				Data:     &twitter.DirectMessageData{Text: "aaaa"},
			},
		},
	})

	fmt.Printf("\n a %+v \n b %+v \n c %+v\n", a, b, c)
	return nil
}
