package main

import (
	"fmt"
	"net/http"

	models "github.com/golangnigeria/MyNneFarm/internal/model"
)

// GetFarms handles the HTTP request to retrieve all farms.
// It uses the NeonDBRepo to fetch the farms and writes the response in JSON format.
func (app *application) GetFarms(w http.ResponseWriter, r *http.Request) {
	// This is a placeholder for the actual implementation of the GetFarms handler.
	farms, err := app.DB.GetFarms()
	if err != nil {
		app.logger.Println("Error retrieving farms:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"farms": farms}, nil)
	if err != nil {
		app.logger.Println("Error writing JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// InsertFarms handles the HTTP request to insert a new farm.
func (app *application) InsertFarms(w http.ResponseWriter, r *http.Request) {
	var farm models.FarmModel

	// Read the incoming JSON into the farm model
	err := app.readJSON(w, r, &farm)
	if err != nil {
		app.logger.Println("Error reading JSON request:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Validate the incoming data
	if err := farm.Validate(); err != nil {
		app.logger.Println("Error validating farm:", err)
		http.Error(w, "Validation Error", http.StatusBadRequest)
		return
	}

	// Save to DB
	farmID, err := app.DB.InsertFarm(farm)
	if err != nil {
		app.logger.Println("Error inserting farm:", err)
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	farm.ID = int(farmID)
	w.Header().Set("Location", fmt.Sprintf("/farms/%d", farm.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"farm": farm}, nil)
	if err != nil {
		app.logger.Println("Error writing JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
