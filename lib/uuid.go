package lib

import uuid "github.com/satori/go.uuid"

type UUIDManager interface {
	GenerateUUIDString() string
}

type uuidManager struct{}

var u *uuidManager

func NewUUIDManager() UUIDManager {
	if u == nil {
		u = &uuidManager{}
	}
	return u
}

func (u *uuidManager) GenerateUUIDString() string {
	return uuid.NewV4().String()
}
