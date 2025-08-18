package repo

import (
	"database/sql"
	"fmt"

	instructor_entity "github.com/SulaimonYNWA/GoTemple/entity/instructors"
)

type InstructorRepo interface {
	GetAll() ([]instructor_entity.Instructor, error)
}

type instructorRepo struct {
	db *sql.DB
}

func NewInstructorRepo(db *sql.DB) InstructorRepo {
	return &instructorRepo{db: db}
}

func (r *instructorRepo) GetAll() ([]instructor_entity.Instructor, error) {
	rows, err := r.db.Query(`
		SELECT id, first_name, last_name, email, phone, department, school_id, created_at
		FROM instructors
	`)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var instructors []instructor_entity.Instructor
	for rows.Next() {
		var ins instructor_entity.Instructor
		if err := rows.Scan(
			&ins.ID,
			&ins.FirstName,
			&ins.LastName,
			&ins.Email,
			&ins.Phone,
			&ins.Department,
			&ins.SchoolID,
			&ins.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		instructors = append(instructors, ins)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return instructors, nil
}

