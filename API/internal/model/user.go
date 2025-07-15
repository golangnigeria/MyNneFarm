package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"` // UUID primary key
	Email     string    `json:"email" db:"email"`
	FullName  string    `json:"full_name" db:"full_name"`
	Phone     string    `json:"phone" db:"phone"`
	Role      string    `json:"role" db:"role"` // "user", "admin"
	Wallet    float64   `json:"wallet" db:"wallet"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
