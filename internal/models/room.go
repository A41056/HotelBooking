package models

import (
	"main.go/internal/const"
	"main.go/internal/domain"
)

// Room model
type Room struct {
	Base
	RoomNumber string        `json:"room_number"`
	Type       string        `json:"type"`
	Price      float64       `json:"price"`
	Status     _const.Status `json:"status"`
}

func (r *Room) ToDomain() *domain.Room {
	return &domain.Room{
		Base: domain.Base{
			ID:         r.Base.ID,
			CreatedAt:  r.Base.CreatedAt,
			ModifiedAt: r.Base.ModifiedAt,
		},
		RoomNumber: r.RoomNumber,
		Type:       r.Type,
		Price:      r.Price,
		Status:     r.Status,
	}
}
