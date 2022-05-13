package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Harvest struct {
	Base
	LinkUUID      string    `json:"link_uuid"`
	RawData       string    `json:"raw_data"`
	PageLink      string    `json:"page_link"`
	Info          string    `json:"info"`
	DisappearedAt time.Time `json:"disappeared_at"`
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
