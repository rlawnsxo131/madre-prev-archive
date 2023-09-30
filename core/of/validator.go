package of

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	_onceValidate      sync.Once
	singletonValidator *validator.Validate
)

func Validator() *validator.Validate {
	_onceValidate.Do(func() {
		singletonValidator = validator.New(validator.WithRequiredStructEnabled())
	})
	return singletonValidator
}
