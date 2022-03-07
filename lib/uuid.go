package lib

import (
	uuid "github.com/satori/go.uuid"
)

var u *uuidManager

type UUIDManager interface {
	GenerateUUIDString() string
}

type uuidManager struct{}

func GetUUIDManager() UUIDManager {
	once.Do(func() {
		u = &uuidManager{}
	})
	return u
}

func (u *uuidManager) GenerateUUIDString() string {
	return uuid.NewV4().String()
}
