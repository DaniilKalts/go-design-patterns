package domain

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID `bson:"_id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`

	OwnerID uuid.UUID `bson:"owner_id"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
