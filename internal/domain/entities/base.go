package entities

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (b *Base) SetDefaultValues() {
	b.ID = primitive.NewObjectIDFromTimestamp(time.Now())
	b.UUID = uuid.New().String()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *Base) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}
