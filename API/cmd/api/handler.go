package main

import (
	"errors"
	"fmt"
	"net/http"

	models "github.com/golangnigeria/MyNneFarm/internal/model"
	"github.com/golangnigeria/MyNneFarm/internal/repository/repo"
	"github.com/golangnigeria/MyNneFarm/internal/validator"
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

// CreateUser handles the HTTP request to create a new user.
func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FullName string `json:"full_name" db:"full_name"`
		Email    string `json:"email" db:"email"`
		Phone    string `json:"phone" db:"phone"`
		Password string `json:"password" db:"password"`
	}

	// Read the incoming JSON into the user model
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.logger.Println("Error reading JSON request:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user := &models.User{
		FullName:  input.FullName,
		Email:     input.Email,
		Phone:     input.Phone,
		Activated: false,              // Default to false
		Roles:     []string{"farmer"}, // Default role
	}

	// Set the password using the Password model
	if err := user.Password.Set(input.Password); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if models.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Save to DB
	err = app.DB.CreateUser(user)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return

	}


	// send a welcome email
	 app.background(func() {
		if err := app.mailer.Send(user.Email, "user_welcome.html", user); err != nil {
			app.logger.Println("Error sending welcome email:", err)
		}
	})
	

	err = app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil)
	if err != nil {
		app.logger.Println("Error writing JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
