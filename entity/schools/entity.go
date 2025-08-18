// entity/school.go
package entity

import "time"

type School struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Address     string    `db:"address"`
	Phone       string    `db:"phone"`
	Email       string    `db:"email"`
	OwnerUserID int       `db:"owner_user_id"`
	CreatedAt   time.Time `db:"created_at"`
}
