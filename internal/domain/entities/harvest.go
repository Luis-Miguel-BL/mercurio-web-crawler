package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

var HarvestCollectionName = "harvest"

type HarvestType string

const HarvestBuilding HarvestType = "building"

type Harvest struct {
	Base          `bson:",inline"`
	LinkUUID      string      `json:"link_uuid" bson:"link_uuid"`
	RawData       string      `json:"raw_data" bson:"raw_data"`
	PageLink      string      `json:"page_link" bson:"page_link"`
	HarvestType   HarvestType `json:"harvest_type" bson:"harvest_type"`
	Info          string      `json:"info" bson:"info"`
	DisappearedAt time.Time   `json:"disappeared_at" bson:"disappeared_at"`
}

func (h *Harvest) Validate() error {
	h.PageLink = strings.Trim(h.PageLink, " ")
	if strings.Contains(h.PageLink, " ") {
		return errors.New("invalid page link")
	}

	if h.RawData != "" {
		var rawData interface{}
		err := json.Unmarshal([]byte(h.RawData), &rawData)
		if err != nil {
			return fmt.Errorf("invalid raw data %v", err)
		}
	}

	if h.Info != "" {
		var info interface{}
		err := json.Unmarshal([]byte(h.Info), &info)
		if err != nil {
			return fmt.Errorf("invalid info %v", err)
		}
	}
	return nil
}

func (h *Harvest) SetDisappeared() {
	h.DisappearedAt = time.Now()
}
