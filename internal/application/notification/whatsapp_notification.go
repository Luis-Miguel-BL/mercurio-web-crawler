package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mercurio-web-scraping/internal/config"
	"net/http"
)

type WhatsappNotification struct {
	MetaAPIVersion string
	PhoneNumberID  string
	BearerToken    string
}

type sendWhatsAppMessage struct {
	MessagingProduct string               `json:"messaging_product"`
	To               string               `json:"to"`
	Type             string               `json:"type"`
	Template         sendWhatsAppTemplate `json:"template"`
}
type sendWhatsAppTemplate struct {
	Name       string                   `json:"name"`
	Language   sendWhatsAppLanguage     `json:"language"`
	Components []sendWhatsAppComponents `json:"components"`
}

type sendWhatsAppLanguage struct {
	Code string `json:"code"`
}
type sendWhatsAppComponents struct {
	Type       string                   `json:"type"`
	Parameters []sendWhatsAppParameters `json:"parameters"`
}
type sendWhatsAppParameters struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewWhatsappNotification(config config.Config) *WhatsappNotification {

	return &WhatsappNotification{
		MetaAPIVersion: config.MetaAPIVersion,
		PhoneNumberID:  config.PhoneNumberID,
		BearerToken:    config.BearerToken,
	}
}

func (n *WhatsappNotification) SendNotification(notification Notification) error {
	wppSendMessage := sendWhatsAppMessage{
		MessagingProduct: "whatsapp",
		// To: notification.Destination,
		To:   "5537998153343",
		Type: "template",
		Template: sendWhatsAppTemplate{
			Name:     "new_harvest",
			Language: sendWhatsAppLanguage{Code: "pt_BR"},
			Components: []sendWhatsAppComponents{
				{
					Type: "body",
					Parameters: []sendWhatsAppParameters{
						{
							Type: "text",
							Text: "link",
						},
					},
				},
			},
		},
	}
	json_data, err := json.Marshal(wppSendMessage)
	if err != nil {
		panic(err)
	}

	// body := ioutil.NopCloser(strings.NewReader(string(json_data)))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", n.MetaAPIVersion, n.PhoneNumberID), bytes.NewReader(json_data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+n.BearerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	fmt.Println(string(data))
	return nil
}
