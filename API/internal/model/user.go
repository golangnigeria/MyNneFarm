package models

import (
	"errors"
	"time"

	"github.com/golangnigeria/MyNneFarm/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            string    `json:"id" db:"id"`
	FullName      string    `json:"full_name" db:"full_name"`
	Email         string    `json:"email" db:"email"`
	Phone         string    `json:"phone" db:"phone"`
	Password      password  `json:"-" db:"-"` // omit in JSON responses
	Roles         []string  `json:"roles" db:"roles"`
	WalletBalance int       `json:"wallet_balance" db:"wallet_balance"`
	Activated     bool      `json:"activated" db:"activated"`
	Version       int       `json:"version" db:"version"`
	ReferredBy    string    `json:"referred_by" db:"referred_by"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type password struct {
	Plaintext *string
	Hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	bcrypted, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.Plaintext = &plaintextPassword
	p.Hash = bcrypted
	return nil
}

func (p *password) Match(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.FullName != "", "full_name", "must be provided")
	v.Check(len(user.FullName) <= 500, "full_name", "must not be more than 500 bytes long")
	// Call the standalone ValidateEmail() helper.
	ValidateEmail(v, user.Email)

	if user.Password.Plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.Plaintext)
	}

	if user.Password.Hash == nil {
		panic("missing password hash")
	}

}
