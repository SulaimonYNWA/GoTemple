package repo

import (
	"database/sql"
	"fmt"

	school_entity "github.com/SulaimonYNWA/GoTemple/entity/schools"
)

type SchoolRepo interface {
	GetAll() ([]school_entity.School, error)
}

type schoolRepo struct {
	db *sql.DB
}

func NewSchoolRepo(db *sql.DB) SchoolRepo {
	return &schoolRepo{db: db}
}
func (r *schoolRepo) GetAll() ([]school_entity.School, error) {
	rows, err := r.db.Query(`
		SELECT id, name, address, phone, email, owner_user_id, created_at 
		FROM schools
	`)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var schools []school_entity.School
	for rows.Next() {
		var s school_entity.School
		if err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Address,
			&s.Phone,
			&s.Email,
			&s.OwnerUserID,
			&s.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		schools = append(schools, s)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return schools, nil
}
