package _const

// Status enum for room status
type Status int

const (
	Empty Status = iota
	Booked
	Repair
)
