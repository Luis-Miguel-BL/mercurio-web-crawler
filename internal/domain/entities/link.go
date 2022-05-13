package entities

import (
	"errors"
	"strings"
	"time"
)

type HarvestType string

const HarvestBuilding HarvestType = "building"

type Link struct {
	Base
	Url                string      `json:"url"`
	Slug               string      `json:"slug"`
	Origin             string      `json:"origin"`
	Description        string      `json:"description"`
	LastVisit          time.Time   `json:"last_visit"`
	TimeoutInSeconds   int64       `json:"timeout_in_seconds"`
	HarvestType        HarvestType `json:"harvest_type"`
	Active             bool        `json:"active"`
	TotalVisits        int64       `json:"total_visits"`
	TotalErrorInVisits int64       `json:"total_error_in_visits"`
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

func (l *Link) SetVisit(hasError bool) {
	l.LastVisit = time.Now()
	l.TotalVisits++
	if hasError {
		l.TotalErrorInVisits++
	}
}

func (l *Link) CreateHarvest(rawData string, pageLink string, info string) Harvest {
	harvest := Harvest{
		LinkUUID: l.UUID,
		RawData:  rawData,
		PageLink: pageLink,
		Info:     info,
	}
	harvest.SetDefaultValues()

	return harvest
}
