// entity/instructors/entity.go
package entity

import "time"

type Instructor struct {
	ID         int       `db:"id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Email      string    `db:"email"`
	Phone      string    `db:"phone"`
	Department string    `db:"department"`
	SchoolID   int       `db:"school_id"`
	CreatedAt  time.Time `db:"created_at"`
}

