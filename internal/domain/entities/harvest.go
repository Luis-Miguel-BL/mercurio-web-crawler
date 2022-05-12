package entities

import "time"

type Harvest struct {
	RawData       string    `json:"raw_data"`
	PageLink      string    `json:"page_link"`
	Info          string    `json:"info"`
	DisappearedAt time.Time `json:"disappeared_at"`
}
