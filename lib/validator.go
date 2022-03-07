package lib

import (
	"github.com/go-playground/validator/v10"
)

var ValidateManager *validator.Validate = validator.New()
