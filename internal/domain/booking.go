package domain

import (
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	Base
	UserID       uuid.UUID
	RoomID       uuid.UUID
	CheckInDate  time.Time
	CheckOutDate time.Time
	Status       int
}
