package services

import (
	"context"

	"dev.longnt1.git/aessment-hotel-booking.git/internal/domain"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/repositories"
	"github.com/google/uuid"
)

type RoomService interface {
	GetRoomByID(ctx context.Context, roomId uuid.UUID) (*domain.Room, error)
	CreateRoom(ctx context.Context, room *domain.Room) error
	UpdateRoom(ctx context.Context, room *domain.Room) error
	DeleteRoom(ctx context.Context, id uuid.UUID) error
	GetRooms(ctx context.Context, filters map[string]interface{}) ([]*domain.Room, error)
}

type roomService struct {
	roomRepo repositories.RoomRepository
}

func NewRoomService(roomRepo repositories.RoomRepository) RoomService {
	return &roomService{roomRepo: roomRepo}
}

func (rs *roomService) GetRoomByID(ctx context.Context, roomId uuid.UUID) (*domain.Room, error) {
	return rs.roomRepo.GetRoomByID(ctx, roomId)
}

func (rs *roomService) CreateRoom(ctx context.Context, room *domain.Room) error {
	return rs.roomRepo.CreateRoom(ctx, room)
}

func (rs *roomService) UpdateRoom(ctx context.Context, room *domain.Room) error {
	return rs.roomRepo.UpdateRoom(ctx, room)
}

func (rs *roomService) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	return rs.roomRepo.DeleteRoom(ctx, id)
}

func (rs *roomService) GetRooms(ctx context.Context, filters map[string]interface{}) ([]*domain.Room, error) {
	return rs.roomRepo.GetRooms(ctx, filters)
}
