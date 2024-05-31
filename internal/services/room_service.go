package services

import (
	"context"
	"github.com/google/uuid"
	"main.go/internal/models"
	"main.go/internal/repositories"
)

type RoomService struct {
	roomRepo repositories.RoomRepository
}

func NewRoomService(roomRepo repositories.RoomRepository) *RoomService {
	return &RoomService{roomRepo: roomRepo}
}

func (rs *RoomService) GetRoomByID(ctx context.Context, roomId uuid.UUID) (*models.Room, error) {
	return rs.roomRepo.GetRoomByID(ctx, roomId)
}

func (rs *RoomService) CreateRoom(ctx context.Context, room *models.Room) error {
	return rs.roomRepo.CreateRoom(ctx, room)
}

func (rs *RoomService) UpdateRoom(ctx context.Context, room *models.Room) error {
	return rs.roomRepo.UpdateRoom(ctx, room)
}

func (rs *RoomService) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	return rs.roomRepo.DeleteRoom(ctx, id)
}

func (rs *RoomService) GetRooms(ctx context.Context, filters map[string]interface{}) ([]models.Room, error) {
	return rs.roomRepo.GetRooms(ctx, filters)
}
