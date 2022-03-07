package lib

import uuid "github.com/satori/go.uuid"

var u *uuidManager

type UUIDManager interface {
	GenerateUUIDString() string
}

type uuidManager struct{}

func NewUUIDManager() UUIDManager {
	if u == nil {
		u = &uuidManager{}
	}
	return u
}

func (u *uuidManager) GenerateUUIDString() string {
	return uuid.NewV4().String()
}
