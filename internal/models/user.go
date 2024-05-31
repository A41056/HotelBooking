package models

// User model
type User struct {
	Base
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	PasswordSalt string    `json:"-"`
	Roles        []*Role   `gorm:"many2many:user_roles;" json:"roles"`
	Bookings     []Booking `json:"bookings"`
}
