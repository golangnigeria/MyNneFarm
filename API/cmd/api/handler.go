package main

import (
	"net/http"
	"time"

	models "github.com/golangnigeria/MyNneFarm/internal/model"
)

func (app *application) GetFarms(w http.ResponseWriter, r *http.Request) {

	// Farm data for demonstration purposes
	farm := models.Farm{
		ID:             1,
		Title:          "Oyo Maize Farm",
		Crop:           "Maize",
		Location:       "Oyo, Nigeria",
		ImageURL:       "http://example.com/image.jpg",
		Description:    "A beautiful maize farm in Oyo.",
		PricePerUnit:   5000.00,
		ExpectedROI:    1.2,
		UnitsAvailable: 100,
		UnitsSold:      50,
		StartDate:      time.Now(),
		HarvestDate:    time.Now().AddDate(0, 0, 30),
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	headers := http.Header{"Custom-Header": []string{"CustomValue"}}

	err := app.writeJSON(w, http.StatusOK, envelope{
		"farm": farm,
	}, headers)
	if err != nil {
		app.logger.Println("Error writing JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
