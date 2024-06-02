package repositories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"main.go/internal/models"
)

type BookingRepository interface {
	Create(booking *models.Booking) error
	Update(booking *models.Booking) error
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (*models.Booking, error)
	GetAll() ([]models.Booking, error)
	GetByUserID(userID uuid.UUID) ([]models.Booking, error)
	GetOverlappingBookings(roomID uuid.UUID, checkInDate, checkOutDate time.Time) ([]models.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) Update(booking *models.Booking) error {
	return r.db.Save(booking).Error
}

func (r *bookingRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Booking{}, id).Error
}

func (r *bookingRepository) GetByID(id uuid.UUID) (*models.Booking, error) {
	var booking models.Booking
	err := r.db.First(&booking, id).Error
	return &booking, err
}

func (r *bookingRepository) GetAll() ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetByUserID(userID uuid.UUID) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetOverlappingBookings(roomID uuid.UUID, checkInDate, checkOutDate time.Time) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Where("room_id = ? AND status = ? AND check_out_date > ? AND check_in_date < ?", roomID, models.Booked, checkInDate, checkOutDate).Find(&bookings).Error
	return bookings, err
}
