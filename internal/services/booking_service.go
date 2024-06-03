package services

import (
	"errors"
	"time"

	_const "dev.longnt1.git/aessment-hotel-booking.git/internal/const"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/domain"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/repositories"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(userID, roomID uuid.UUID, checkInDate, checkOutDate time.Time, status _const.Status) (*domain.Booking, error)
	UpdateBooking(id uuid.UUID, checkInDate, checkOutDate time.Time, status _const.Status) (*domain.Booking, error)
	DeleteBooking(id uuid.UUID) error
	GetBookingByID(id uuid.UUID) (*domain.Booking, error)
	GetAllBookings() ([]domain.Booking, error)
	GetBookingsByUserID(userID uuid.UUID) ([]domain.Booking, error)
}

type bookingService struct {
	repo repositories.BookingRepository
}

func NewBookingService(repo repositories.BookingRepository) BookingService {
	return &bookingService{repo: repo}
}

func (s *bookingService) CreateBooking(userID, roomID uuid.UUID, checkInDate, checkOutDate time.Time, status _const.Status) (*domain.Booking, error) {
	if status != _const.Booked && status != _const.Cancelled {
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

	booking := &domain.Booking{
		UserID:       userID,
		RoomID:       roomID,
		CheckInDate:  checkInDate,
		CheckOutDate: checkOutDate,
		Status:       int(status),
	}
	if err := s.repo.Create(booking); err != nil {
		return nil, err
	}
	return booking, nil
}

func (s *bookingService) UpdateBooking(id uuid.UUID, checkInDate, checkOutDate time.Time, status _const.Status) (*domain.Booking, error) {
	if status != _const.Booked && status != _const.Cancelled {
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
	booking.Status = int(status)

	if err := s.repo.Update(booking); err != nil {
		return nil, err
	}
	return booking, nil
}

func (s *bookingService) DeleteBooking(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *bookingService) GetBookingByID(id uuid.UUID) (*domain.Booking, error) {
	return s.repo.GetByID(id)
}

func (s *bookingService) GetAllBookings() ([]domain.Booking, error) {
	return s.repo.GetAll()
}

func (s *bookingService) GetBookingsByUserID(userID uuid.UUID) ([]domain.Booking, error) {
	return s.repo.GetByUserID(userID)
}
