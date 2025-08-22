// entity/instructors/entity.go
package entity

import (
	"time"
)

type Instructor struct {
	ID             int       `db:"id"`
	Name           string    `db:"name"`
	Phone          string    `db:"phone"`
	Email          string    `db:"email"`
	UserID         int64     `db:"user_id"`
	SchoolName     string    `db:"name"`
	SchoolID       int       `db:"school_id"`
	Specialization string    `db:"specialization"`
	Salary         float64   `db:"salary"`
	CreatedAt      time.Time `db:"created_at"`
}
