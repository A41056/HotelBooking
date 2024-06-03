package domain

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	ModifiedAt time.Time
}
