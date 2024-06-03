package controllers

import (
	"net/http"

	_const "dev.longnt1.git/aessment-hotel-booking.git/internal/const"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/domain"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoomController struct {
	roomService services.RoomService
}

func NewRoomController(roomService services.RoomService) *RoomController {
	return &RoomController{roomService: roomService}
}

func (rc *RoomController) GetRoomByID(c *gin.Context) {
	roomIdStr := c.Query("id")
	if roomIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrRequireRoom})
		return
	}

	roomId, err := uuid.Parse(roomIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidRoomID})
		return
	}

	if roomId == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidRoomID})
		return
	}

	ctx := c.Request.Context()
	room, err := rc.roomService.GetRoomByID(ctx, roomId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (rc *RoomController) CreateRoom(c *gin.Context) {
	var room domain.Room
	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrFailedDecodeRequestBody})
		return
	}

	ctx := c.Request.Context()
	if err := rc.roomService.CreateRoom(ctx, &room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, room)
}

func (rc *RoomController) UpdateRoom(c *gin.Context) {
	var room domain.Room
	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrFailedDecodeRequestBody})
		return
	}

	ctx := c.Request.Context()
	if err := rc.roomService.UpdateRoom(ctx, &room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (rc *RoomController) DeleteRoom(c *gin.Context) {
	roomIdStr := c.Query("id")
	if roomIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrRequireRoom})
		return
	}

	roomId, err := uuid.Parse(roomIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrInvalidRoomID})
		return
	}

	ctx := c.Request.Context()
	if err := rc.roomService.DeleteRoom(ctx, roomId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}

func (rc *RoomController) GetRoomsByFilter(c *gin.Context) {
	var filters map[string]interface{}
	if err := c.BindJSON(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _const.ErrFailedDecodeRequestBody})
		return
	}

	ctx := c.Request.Context()
	rooms, err := rc.roomService.GetRooms(ctx, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rooms)
}
