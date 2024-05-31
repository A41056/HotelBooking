package models

import (
	"time"

	"github.com/google/uuid"
)

// Booking model
type Booking struct {
	Base
	UserID       uuid.UUID `json:"user_id"`
	RoomID       uuid.UUID `json:"room_id"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	Status       int       `json:"status"`
}
