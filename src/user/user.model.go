package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID          uint
	Email       string
	UserName    sql.NullString
	DisplayName string
	PhotoUrl    sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
