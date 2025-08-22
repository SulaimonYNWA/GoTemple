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
		SELECT 
			i.id,
			i.school_id,
			i.specialization,
			i.salary,
			u.id,
			u.name,
			u.email,
			u.phone,
			u.created_at,
			s.name
		FROM instructors i
		JOIN users u ON i.user_id = u.id
		Join schools s ON i.school_id = s.id
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
			&ins.SchoolID,
			&ins.SchoolName,
			&ins.Specialization,
			&ins.Salary,
			&ins.Name,
			&ins.Email,
			&ins.Phone,
			&ins.CreatedAt,
			&ins.SchoolName,
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
