package repositories

import (
	"time"

	_const "dev.longnt1.git/aessment-hotel-booking.git/internal/const"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/domain"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(booking *domain.Booking) error
	Update(booking *domain.Booking) error
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (*domain.Booking, error)
	GetAll() ([]domain.Booking, error)
	GetByUserID(userID uuid.UUID) ([]domain.Booking, error)
	GetOverlappingBookings(roomID uuid.UUID, checkInDate, checkOutDate time.Time) ([]domain.Booking, error)
}

func toDomainBase(model *models.Base) *domain.Base {
	return &domain.Base{
		ID:         model.ID,
		CreatedAt:  model.CreatedAt,
		ModifiedAt: model.ModifiedAt,
	}
}

func toModelBase(domain *domain.Base) *models.Base {
	return &models.Base{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		ModifiedAt: domain.ModifiedAt,
	}
}

func toDomainBooking(model *models.Booking) *domain.Booking {
	return &domain.Booking{
		Base:         *toDomainBase(&model.Base),
		UserID:       model.UserID,
		RoomID:       model.RoomID,
		CheckInDate:  model.CheckInDate,
		CheckOutDate: model.CheckOutDate,
		Status:       model.Status,
	}
}

func toModelBooking(domain *domain.Booking) *models.Booking {
	return &models.Booking{
		Base:         *toModelBase(&domain.Base),
		UserID:       domain.UserID,
		RoomID:       domain.RoomID,
		CheckInDate:  domain.CheckInDate,
		CheckOutDate: domain.CheckOutDate,
		Status:       domain.Status,
	}
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(booking *domain.Booking) error {
	modelBooking := toModelBooking(booking)
	return r.db.Create(modelBooking).Error
}

func (r *bookingRepository) Update(booking *domain.Booking) error {
	modelBooking := toModelBooking(booking)
	return r.db.Save(modelBooking).Error
}

func (r *bookingRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Booking{}, id).Error
}

func (r *bookingRepository) GetByID(id uuid.UUID) (*domain.Booking, error) {
	var modelBooking models.Booking
	err := r.db.First(&modelBooking, id).Error
	if err != nil {
		return nil, err
	}
	return toDomainBooking(&modelBooking), nil
}

func (r *bookingRepository) GetAll() ([]domain.Booking, error) {
	var modelBookings []models.Booking
	err := r.db.Find(&modelBookings).Error
	if err != nil {
		return nil, err
	}
	var domainBookings []domain.Booking
	for _, modelBooking := range modelBookings {
		domainBookings = append(domainBookings, *toDomainBooking(&modelBooking))
	}
	return domainBookings, nil
}

func (r *bookingRepository) GetByUserID(userID uuid.UUID) ([]domain.Booking, error) {
	var modelBookings []models.Booking
	err := r.db.Where("user_id = ?", userID).Find(&modelBookings).Error
	if err != nil {
		return nil, err
	}
	var domainBookings []domain.Booking
	for _, modelBooking := range modelBookings {
		domainBookings = append(domainBookings, *toDomainBooking(&modelBooking))
	}
	return domainBookings, nil
}

func (r *bookingRepository) GetOverlappingBookings(roomID uuid.UUID, checkInDate, checkOutDate time.Time) ([]domain.Booking, error) {
	var modelBookings []models.Booking
	err := r.db.Where("room_id = ? AND status = ? AND check_out_date > ? AND check_in_date < ?", roomID, _const.Booked, checkInDate, checkOutDate).Find(&modelBookings).Error
	if err != nil {
		return nil, err
	}
	var domainBookings []domain.Booking
	for _, modelBooking := range modelBookings {
		domainBookings = append(domainBookings, *toDomainBooking(&modelBooking))
	}
	return domainBookings, nil
}
