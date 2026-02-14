package data

import (
	"time"

	"github.com/andreshungbz/lab3-json-handling/internal/validator"
)

// Room maps a hotel room entity.
type Room struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"-"`
	RoomNumber   int32     `json:"room_number"`
	RoomType     string    `json:"room_type"`
	MaxOccupancy int32     `json:"max_occupancy"`
	HasBalcony   bool      `json:"has_balcony"`
	Available    bool      `json:"available"`
}

// ValidateRoom does simple validation to ensure RoomNumber and MaxOccupancy
// are positive values.
func ValidateRoom(v *validator.Validator, room *Room) {
	v.Check(room.RoomNumber > 0, "room_number", "Room number must be a positive number")
	v.Check(room.MaxOccupancy > 0, "max_occupancy", "Max occupancy must be a positive number")
}
