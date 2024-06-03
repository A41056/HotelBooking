package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main.go/internal/const"
	"main.go/internal/services"
)

type BookingController struct {
	bookingService services.BookingService
}

func NewBookingController(bookingService services.BookingService) *BookingController {
	return &BookingController{bookingService: bookingService}
}

func (bc *BookingController) CreateBooking(c *gin.Context) {
	var req struct {
		UserID       string `json:"user_id"`
		RoomID       string `json:"room_id"`
		CheckInDate  string `json:"check_in_date"`
		CheckOutDate string `json:"check_out_date"`
		Status       int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidUserID})
		return
	}

	roomID, err := uuid.Parse(req.RoomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidRoomID})
		return
	}

	checkInDate, err := time.Parse(time.RFC3339, req.CheckInDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidCheckInDate})
		return
	}

	checkOutDate, err := time.Parse(time.RFC3339, req.CheckOutDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidCheckOutDate})
		return
	}

	booking, err := bc.bookingService.CreateBooking(userID, roomID, checkInDate, checkOutDate, _const.Status(req.Status))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

func (bc *BookingController) UpdateBooking(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidBookingID})
		return
	}

	var req struct {
		CheckInDate  string `json:"check_in_date"`
		CheckOutDate string `json:"check_out_date"`
		Status       int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkInDate, err := time.Parse(time.RFC3339, req.CheckInDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidCheckInDate})
		return
	}

	checkOutDate, err := time.Parse(time.RFC3339, req.CheckOutDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidCheckOutDate})
		return
	}

	booking, err := bc.bookingService.UpdateBooking(id, checkInDate, checkOutDate, _const.Status(req.Status))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func (bc *BookingController) DeleteBooking(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidBookingID})
		return
	}

	if err := bc.bookingService.DeleteBooking(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (bc *BookingController) GetBookingByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidBookingID})
		return
	}

	booking, err := bc.bookingService.GetBookingByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func (bc *BookingController) GetAllBookings(c *gin.Context) {
	bookings, err := bc.bookingService.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func (bc *BookingController) GetBookingsByUserID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidUserID})
		return
	}

	bookings, err := bc.bookingService.GetBookingsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}
