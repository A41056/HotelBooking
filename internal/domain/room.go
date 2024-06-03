package domain

import _const "dev.longnt1.git/aessment-hotel-booking.git/internal/const"

type Room struct {
	Base
	RoomNumber string
	Type       string
	Price      float64
	Status     _const.Status
}
