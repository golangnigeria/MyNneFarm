package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	models "github.com/golangnigeria/MyNneFarm/internal/model"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrUserNotFound   = errors.New("user not found")
)

func (m *NeonDBRepo) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		INSERT INTO users (
			 full_name, email, phone, password_hash, roles,
			wallet_balance, activated, version, referred_by,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11
		) RETURNING id, version, created_at
	`

	// Ensure timestamps are set
	now := time.Now().UTC()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = now
	}

	err := m.DB.QueryRowContext(ctx, query,
		user.FullName,
		user.Email,
		user.Phone,
		user.Password.Hash,
		user.Roles,
		user.WalletBalance,
		user.Activated,
		user.Version,
		user.ReferredBy,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID, &user.Version, &user.CreatedAt)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (m *NeonDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, created_at, email, password, roles,
		wallet_balance, activated
		FROM users
		WHERE email = $1`

	var user models.User

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Email,
		&user.Password,
		&user.Roles,
		&user.WalletBalance,
		&user.Activated,
	)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrUserNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

// UpdateUser updates an existing user in the database.
func (m *NeonDBRepo) UpdateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
			UPDATE users
				SET full_name = $1, email = $2, password_hash = $3, activated = $4, version = version + 1
			WHERE id = $5 AND version = $6
			RETURNING version`

	 

	err := m.DB.QueryRowContext(ctx, query,
		user.FullName,
		user.Email,
		user.Password.Hash,
		user.Activated,
		user.ID,
		user.Version,
	).Scan(&user.Version)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return ErrUserNotFound
		default:
			return err
		}
	}

	return nil
}
