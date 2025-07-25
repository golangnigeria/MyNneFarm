package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int       `json:"id" db:"id"`
	FullName      string    `json:"full_name" db:"full_name"`
	Email         string    `json:"email" db:"email"`
	Phone         string    `json:"phone" db:"phone"`
	Password      string    `json:"-" db:"password"` // stored as hashed password, omitted in JSON
	Roles         []string  `json:"roles,omitempty" db:"roles"`
	WalletBalance int       `json:"wallet_balance" db:"wallet_balance"`
	Activated     bool      `json:"activated" db:"activated"`
	Version       int       `json:"version" db:"version"`
	ReferredBy    string    `json:"referred_by,omitempty" db:"referred_by"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}



func (u *User) PasswordMatches (plainText string) (bool, error){
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil{
		switch{
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return  false, nil
		}
	}

	return true, nil
}