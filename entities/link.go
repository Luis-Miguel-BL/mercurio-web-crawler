package entities

import "time"

type HarvestType string

const HarvestBuilding HarvestType = "building"

type Link struct {
	Base
	Url              string      `json:"url"`
	Slug             string      `json:"slug"`
	Origin           string      `json:"origin"`
	Description      string      `json:"description"`
	LastVisit        time.Time   `json:"last_visit"`
	TimeoutInSeconds int64       `json:"timeout_in_seconds"`
	HarvestType      HarvestType `json:"harvest_type"`
	Active           bool        `json:"active"`
}
