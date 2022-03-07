package lib

import (
	uuid "github.com/satori/go.uuid"
)

type uuidManager struct{}

func NewUUIDManager() *uuidManager {
	return &uuidManager{}
}

func (u *uuidManager) GenerateUUIDString() string {
	return uuid.NewV4().String()
}
