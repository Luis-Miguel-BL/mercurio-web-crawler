package entities

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Base) SetDefaultValues() {
	b.UUID = uuid.New().String()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *Base) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}
