package domain

import _const "main.go/internal/const"

type Room struct {
	Base
	RoomNumber string
	Type       string
	Price      float64
	Status     _const.Status
}
