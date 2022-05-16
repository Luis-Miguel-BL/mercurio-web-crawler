package entities

import (
	"errors"
	"strings"
	"time"
)

var LinkCollectionName = "link"

type Link struct {
	Base               `bson:",inline"`
	Url                string    `json:"url" bson:"url"`
	Slug               string    `json:"slug" bson:"slug"`
	Origin             string    `json:"origin" bson:"origin"`
	Description        string    `json:"description" bson:"description"`
	LastVisit          time.Time `json:"last_visit" bson:"last_visit"`
	TimeoutInSeconds   int64     `json:"timeout_in_seconds" bson:"timeout_in_seconds"`
	Active             bool      `json:"active" bson:"active"`
	TotalVisits        int64     `json:"total_visits" bson:"total_visits"`
	TotalErrorInVisits int64     `json:"total_error_in_visits" bson:"total_error_in_visits"`
}

func (l *Link) Validate() (err error) {
	l.Slug = strings.Trim(l.Slug, " ")
	if strings.Contains(l.Slug, " ") {
		return errors.New("invalid slug")
	}

	l.Url = strings.Trim(l.Url, " ")
	if strings.Contains(l.Url, " ") {
		return errors.New("invalid url")
	}
	return nil

}

func (l *Link) SetVisit() {
	l.LastVisit = time.Now()
	l.TotalVisits++

}
func (l *Link) SetErrorVisit() {
	l.SetVisit()
	l.TotalErrorInVisits++
}

func (l *Link) CreateHarvest(rawData string, pageLink string, info string, harvestType HarvestType) Harvest {
	harvest := Harvest{
		LinkUUID:    l.UUID,
		RawData:     rawData,
		PageLink:    pageLink,
		HarvestType: harvestType,
		Info:        info,
	}
	harvest.SetDefaultValues()

	return harvest
}
