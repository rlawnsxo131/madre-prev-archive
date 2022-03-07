package lib

import (
	uuid "github.com/satori/go.uuid"
)

type Utils interface {
	GenerateUUID() string
	IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int
}

type utils struct{}

func NewUtils() Utils {
	return &utils{}
}

// uuid
func (u *utils) GenerateUUID() string {
	uuid := uuid.NewV4()
	return uuid.String()
}

// default value
func (u *utils) IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int {
	if value == 0 {
		value = defaultValue
	}
	return value
}
