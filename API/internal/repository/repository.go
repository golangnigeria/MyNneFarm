package repository

import (
	"database/sql"

	models "github.com/golangnigeria/MyNneFarm/internal/model"
)

type DatabaseRepository interface {
	Connection() *sql.DB
	GetFarms() ([]models.FarmModel, error)
	InsertFarm(farm models.FarmModel) (int64, error)
}
