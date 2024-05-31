package repositories

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"main.go/internal/models"
)

type RoomRepository interface {
	GetRoomByID(ctx context.Context, roomId uuid.UUID) (*models.Room, error)
	CreateRoom(ctx context.Context, room *models.Room) error
	UpdateRoom(ctx context.Context, room *models.Room) error
	DeleteRoom(ctx context.Context, roomId uuid.UUID) error
	GetRooms(ctx context.Context, filters map[string]interface{}) ([]models.Room, error)
}

type roomRepo struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepo{DB: db}
}

func (r *roomRepo) GetRoomByID(ctx context.Context, roomId uuid.UUID) (*models.Room, error) {
	var room models.Room
	err := r.DB.Where("id = ?", roomId).First(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepo) CreateRoom(ctx context.Context, room *models.Room) error {
	return r.DB.Create(room).Error
}

func (r *roomRepo) UpdateRoom(ctx context.Context, room *models.Room) error {
	return r.DB.Save(room).Error
}

func (r *roomRepo) DeleteRoom(ctx context.Context, roomId uuid.UUID) error {
	return r.DB.Where("id = ?", roomId).Delete(&models.Room{}).Error
}

func (r *roomRepo) GetRooms(ctx context.Context, filters map[string]interface{}) ([]models.Room, error) {
	var rooms []models.Room

	query := r.DB
	for key, value := range filters {
		switch key {
		case "room_number":
			query = query.Where("room_number = ?", value)
		case "status":
			query = query.Where("status = ?", value)
		case "price_type":
			query = query.Where("price_type = ?", value)
		}
	}

	if err := query.Find(&rooms).Error; err != nil {
		return nil, err
	}

	return rooms, nil
}
