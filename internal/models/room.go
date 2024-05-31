package models

import (
	"main.go/internal/const"
)

// Room model
type Room struct {
	Base
	RoomNumber string        `json:"room_number"`
	Type       string        `json:"type"`
	Price      float64       `json:"price"`
	Status     _const.Status `json:"status"`
}
