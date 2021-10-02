package user

import (
	"time"
)

type UserService interface {
	FindById(id int) *User
}

func FindById(id int) *User {
	user := &User{
		ID:          1,
		Email:       "juntae@gmail.com",
		DisplayName: "juntae",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return user
}
