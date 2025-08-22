// entity/courses/entity.go
package entity

import "time"

type Course struct {
	ID          int       `db:"id"`
	SchoolID    int       `db:"school_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Credits     int       `db:"credits"`
	CreatedAt   time.Time `db:"created_at"`
}

