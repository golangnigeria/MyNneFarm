package repo

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	models "github.com/golangnigeria/MyNneFarm/internal/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrUserNotFound   = errors.New("user not found")
)

func (m *NeonDBRepo) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Hash password before inserting
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO users (
			full_name, email, phone, password, roles,
			wallet_balance, activated, version, referred_by,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11
		) RETURNING id, version, created_at
	`

	now := time.Now().UTC()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = now
	}

	err = m.DB.QueryRowContext(ctx, query,
		user.FullName,
		user.Email,
		user.Phone,
		hashedPassword, // store hash, not plain password
		user.Roles,
		user.WalletBalance,
		user.Activated,
		user.Version,
		user.ReferredBy,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID, &user.Version, &user.CreatedAt)

	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_email_key"`) {
			return ErrDuplicateEmail
		}
		return err
	}

	return nil
}

func (m *NeonDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, full_name, email, phone, password, roles,
		       wallet_balance, activated, version, referred_by,
		       created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user models.User

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Phone,
		&user.Password, // this is the stored hash
		&user.Roles,
		&user.WalletBalance,
		&user.Activated,
		&user.Version,
		&user.ReferredBy,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
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
		user.Password,
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
