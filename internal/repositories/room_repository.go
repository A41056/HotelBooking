package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"main.go/internal/domain"
	"main.go/internal/models"
)

type RoomRepository interface {
	GetRoomByID(ctx context.Context, roomId uuid.UUID) (*domain.Room, error)
	CreateRoom(ctx context.Context, room *domain.Room) error
	UpdateRoom(ctx context.Context, room *domain.Room) error
	DeleteRoom(ctx context.Context, roomId uuid.UUID) error
	GetRooms(ctx context.Context, filters map[string]interface{}) ([]*domain.Room, error)
}

type roomRepo struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepo{DB: db}
}

func (r *roomRepo) GetRoomByID(ctx context.Context, roomId uuid.UUID) (*domain.Room, error) {
	var room models.Room
	err := r.DB.Where("id = ?", roomId).First(&room).Error
	if err != nil {
		return nil, err
	}

	return room.ToDomain(), nil
}

func (r *roomRepo) CreateRoom(ctx context.Context, room *domain.Room) error {
	return r.DB.Create(room).Error
}

func (r *roomRepo) UpdateRoom(ctx context.Context, room *domain.Room) error {
	return r.DB.Save(room).Error
}

func (r *roomRepo) DeleteRoom(ctx context.Context, roomId uuid.UUID) error {
	return r.DB.Where("id = ?", roomId).Delete(&models.Room{}).Error
}

func (r *roomRepo) GetRooms(ctx context.Context, filters map[string]interface{}) ([]*domain.Room, error) {
	var roomModels []models.Room

	query := r.DB
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	if err := query.Find(&roomModels).Error; err != nil {
		return nil, err
	}

	rooms := make([]*domain.Room, 0)
	for _, room := range roomModels {
		rooms = append(rooms, room.ToDomain())
	}

	return rooms, nil
}
