package entity

import (
	"time"
	"user-service/pkg/enum"
)

type User struct {
	Id        string      `db:"id"`
	FirstName string      `db:"first_name"`
	LastName  string      `db:"last_name"`
	Phone     string      `db:"phone"`
	Email     string      `db:"email"`
	Role      enum.ROLE   `db:"role"`
	Status    enum.STATUS `db:"status"`
	Password  string      `db:"password"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
}
