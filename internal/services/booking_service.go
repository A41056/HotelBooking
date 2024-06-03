package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"main.go/internal/const"
	"main.go/internal/models"
	"main.go/internal/repositories"
)

type BookingService interface {
	CreateBooking(userID, roomID uuid.UUID, checkInDate, checkOutDate time.Time, status int) (*models.Booking, error)
	UpdateBooking(id uuid.UUID, checkInDate, checkOutDate time.Time, status int) (*models.Booking, error)
	DeleteBooking(id uuid.UUID) error
	GetBookingByID(id uuid.UUID) (*models.Booking, error)
	GetAllBookings() ([]models.Booking, error)
	GetBookingsByUserID(userID uuid.UUID) ([]models.Booking, error)
}

type bookingService struct {
	repo repositories.BookingRepository
}

func NewBookingService(repo repositories.BookingRepository) BookingService {
	return &bookingService{repo: repo}
}

func (s *bookingService) CreateBooking(userID, roomID uuid.UUID, checkInDate, checkOutDate time.Time, status int) (*models.Booking, error) {
	if status != models.Booked && status != models.Cancelled {
		return nil, errors.New(_const.ErrInvalidStatus)
	}

	if time.Now().After(checkInDate) || time.Now().After(checkOutDate) {
		return nil, errors.New(_const.ErrPastDateBooking)
	}

	// Check for overlapping bookings
	overlappingBookings, err := s.repo.GetOverlappingBookings(roomID, checkInDate, checkOutDate)
	if err != nil {
		return nil, err
	}
	if len(overlappingBookings) > 0 {
		return nil, errors.New(_const.ErrOverlapBooking)
	}

	booking := &models.Booking{
		UserID:       userID,
		RoomID:       roomID,
		CheckInDate:  checkInDate,
		CheckOutDate: checkOutDate,
		Status:       status,
	}
	if err := s.repo.Create(booking); err != nil {
		return nil, err
	}
	return booking, nil
}

func (s *bookingService) UpdateBooking(id uuid.UUID, checkInDate, checkOutDate time.Time, status int) (*models.Booking, error) {
	if status != models.Booked && status != models.Cancelled {
		return nil, errors.New(_const.ErrInvalidStatus)
	}

	if time.Now().After(checkInDate) || time.Now().After(checkOutDate) {
		return nil, errors.New(_const.ErrPastDateBooking)
	}

	booking, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if booking == nil {
		return nil, errors.New(_const.ErrBookingNotFound)
	}

	// Check for overlapping bookings
	overlappingBookings, err := s.repo.GetOverlappingBookings(booking.RoomID, checkInDate, checkOutDate)
	if err != nil {
		return nil, err
	}
	if len(overlappingBookings) > 0 {
		return nil, errors.New(_const.ErrOverlapBooking)
	}

	booking.CheckInDate = checkInDate
	booking.CheckOutDate = checkOutDate
	booking.Status = status

	if err := s.repo.Update(booking); err != nil {
		return nil, err
	}
	return booking, nil
}

func (s *bookingService) DeleteBooking(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *bookingService) GetBookingByID(id uuid.UUID) (*models.Booking, error) {
	return s.repo.GetByID(id)
}

func (s *bookingService) GetAllBookings() ([]models.Booking, error) {
	return s.repo.GetAll()
}

func (s *bookingService) GetBookingsByUserID(userID uuid.UUID) ([]models.Booking, error) {
	return s.repo.GetByUserID(userID)
}
