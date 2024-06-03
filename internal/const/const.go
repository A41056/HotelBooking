package _const

// Status enum for room status
type Status int

const (
	Empty Status = iota
	Booked
	Repair

	ErrInvalidUserID           = "invalid user_id"
	ErrInvalidRoomID           = "invalid room_id"
	ErrInvalidCheckInDate      = "invalid check_in_date"
	ErrInvalidCheckOutDate     = "invalid check_out_date"
	ErrInvalidBookingID        = "invalid booking id"
	ErrInvalidUserIDParam      = "invalid user id"
	ErrRequireRoom             = "room is required"
	ErrFailedDecodeRequestBody = "failed to decode request body"
	ErrInvalidEmaiOrPassword   = "invalid email or password"
	ErrUserNotFound            = "user not found"
	ErrAuthorizeHeaderMissing  = "authorization header is missing"
	ErrInvalidAuthHeaderFormat = "invalid authorization header format"
	ErrMissingToken            = "token is missing"
	ErrTooManyRequest          = "too many requests"
	ErrOverlapBooking          = "overlapping bookings are not allowed"
	ErrPastDateBooking         = "bookings can only be made for future dates"
	ErrInvalidStatus           = "invalid booking status"
	ErrBookingNotFound         = "booking not found"
	ErrPasswordTooShort        = "password must be at least 6 characters long"
	ErrUserNameRequire         = "username is required"
	ErrInvalidEmailFormat      = "invalid email format"
	ErrIncorrectPassword       = "incorrect password"
	ErrInvalidTokenSignature   = "invalid token signature"
	ErrInvalidToken            = "invalid token"
)
