package models

import (
	"time"
)

// Farm represents a farm available for investment
type Farm struct {
	ID              int       `json:"id" db:"id"`                         // Auto-increment primary key
	Title           string    `json:"title" db:"title"`                   // e.g. "Oyo Maize Farm"
	Crop            string    `json:"crop" db:"crop"`                     // e.g. "Maize", "Cassava"
	Location        string    `json:"location" db:"location"`             // e.g. "Oyo, Nigeria"
	ImageURL        string    `json:"image_url" db:"image_url"`           // Cloud image for display
	Description     string    `json:"description" db:"description"`       // Full description for display
	PricePerUnit    float64   `json:"price_per_unit" db:"price_per_unit"` // e.g. 5000.00 NGN
	ExpectedROI     float64   `json:"expected_roi" db:"expected_roi"`     // e.g. 1.2 = 20% profit
	UnitsAvailable  int       `json:"units_available" db:"units_available"`
	UnitsSold       int       `json:"units_sold" db:"units_sold"`
	StartDate       time.Time `json:"start_date" db:"start_date"`         // Date the farm starts
	HarvestDate     time.Time `json:"harvest_date" db:"harvest_date"`     // Date users receive Food Credits
	IsActive        bool      `json:"is_active" db:"is_active"`           // Whether it's open for investment
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}


// FarmSummary provides a brief overview of a farm for listings
type FarmSummary struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Crop       string  `json:"crop"`
	Location   string  `json:"location"`
	ImageURL   string  `json:"image_url"`
	PricePerUnit float64 `json:"price_per_unit"`
	ExpectedROI  float64 `json:"expected_roi"`
}