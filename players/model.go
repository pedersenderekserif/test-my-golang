package players

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	ID        uuid.UUID  `json:"id"`
	TeamID    uuid.UUID  `json:"teamId"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type Players []Player
