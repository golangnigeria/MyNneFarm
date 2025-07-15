package repo

import (
	"context"
	"database/sql"
	"time"

	models "github.com/golangnigeria/MyNneFarm/internal/model"
)

type NeonDBRepo struct {
	DB *sql.DB
}

const dbTimeout = 5 * time.Second

func (m *NeonDBRepo) Connection() *sql.DB {
	return m.DB
}
	


// GetFarms retrieves all farms from the database.
func (m *NeonDBRepo) GetFarms() ([]models.FarmModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := ` SELECT 
					id, title, crop, location, coalesce(image_url, '') AS image_url, description, price_per_unit, expected_roi, units_available, units_sold, start_date, harvest_date, is_active, created_at, updated_at 
				FROM 
				  	farms
				ORDER BY created_at DESC;`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var farms []models.FarmModel

	for rows.Next() {
		var farm models.FarmModel
		if err := rows.Scan(
			&farm.ID,
			&farm.Title,
			&farm.Crop,
			&farm.Location,
			&farm.ImageURL,
			&farm.Description,
			&farm.PricePerUnit,
			&farm.ExpectedROI,
			&farm.UnitsAvailable,
			&farm.UnitsSold,
			&farm.StartDate,
			&farm.HarvestDate,
			&farm.IsActive,
			&farm.CreatedAt,
			&farm.UpdatedAt); err != nil {
			return nil, err
		}
		farms = append(farms, farm)
	}

	return farms, nil
}
