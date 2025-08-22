package repo

import (
	"database/sql"
	"fmt"

	course_entity "github.com/SulaimonYNWA/GoTemple/entity/courses"
)

type CourseRepo interface {
	GetBySchoolID(schoolID int) ([]course_entity.Course, error)
}

type courseRepo struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) CourseRepo {
	return &courseRepo{db: db}
}

func (r *courseRepo) GetBySchoolID(schoolID int) ([]course_entity.Course, error) {
	rows, err := r.db.Query(`
		SELECT id, school_id, name, description, credits, created_at
		FROM courses
		WHERE school_id = ?
	`, schoolID)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var courses []course_entity.Course
	for rows.Next() {
		var c course_entity.Course
		if err := rows.Scan(
			&c.ID,
			&c.SchoolID,
			&c.Name,
			&c.Description,
			&c.Credits,
			&c.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		courses = append(courses, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return courses, nil
}

