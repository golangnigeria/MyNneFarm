package models

import "time"

const (
	// TokenTypeActivation is used for account activation tokens.
	TokenTypeActivation = "activation"
)

type Token struct {
	PlainText string
	Hash      []byte
	UserID    string
	Expiry    time.Time
	Scope     string // e.g., "activation", "password_reset"
}

func generateToken(userID string, ttl time.Duration, scope string) (*Token, error) {
	// instance of the Token struct
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}
	return token, nil
}
