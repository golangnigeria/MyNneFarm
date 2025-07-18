package repo

import (
	"database/sql"
)

type NeonDBRepo struct {
	DB *sql.DB
}

func (m *NeonDBRepo) Connection() *sql.DB {
	return m.DB
}
