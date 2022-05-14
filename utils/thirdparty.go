package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

func GenerateUUIDString() string {
	return uuid.NewString()
}
