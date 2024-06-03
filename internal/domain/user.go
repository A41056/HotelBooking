package domain

type User struct {
	Base
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	PasswordSalt string    `json:"-"`
	Roles        []*Role   `json:"roles"`
	Bookings     []Booking `json:"bookings"`
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
