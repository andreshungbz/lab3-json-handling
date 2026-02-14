package main

import (
	"net/http"

	"github.com/andreshungbz/lab3-json-handling/internal/data"
	"github.com/andreshungbz/lab3-json-handling/internal/validator"
)

// showRoomHandler writes JSON of a single hotel room to the HTTP response.
func (app *application) showRoomHandler(w http.ResponseWriter, r *http.Request) {
	// simulate retrieving a hotel room from the database
	room := &data.Room{
		ID:           2,
		RoomNumber:   267,
		RoomType:     "double-bed-room",
		MaxOccupancy: 4,
		HasBalcony:   false,
		Available:    true,
	}

	// use writeJSON for the HTTP response
	err := app.writeJSON(w, http.StatusOK, envelope{"room": room}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// createRoomHandler reads JSON for a hotel room and writes it back to the client.
func (app *application) createRoomHandler(w http.ResponseWriter, r *http.Request) {
	// use an intermediate struct to check validity of JSON
	var input struct {
		RoomNumber   int32  `json:"room_number"`
		RoomType     string `json:"room_type"`
		MaxOccupancy int32  `json:"max_occupancy"`
		HasBalcony   bool   `json:"has_balcony"`
		Available    bool   `json:"available"`
	}

	// attempt readJSON
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// create a Room from the input values
	room := &data.Room{
		ID:           1,
		RoomNumber:   input.RoomNumber,
		RoomType:     input.RoomType,
		MaxOccupancy: input.MaxOccupancy,
		HasBalcony:   input.HasBalcony,
		Available:    input.Available,
	}

	// validate JSON values
	v := validator.New()
	if data.ValidateRoom(v, room); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// use writeJSON for the HTTP response
	err = app.writeJSON(w, http.StatusCreated, envelope{"room": room}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
