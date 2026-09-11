package request

import (
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) (err error) {
	validate := validator.New()
	err = validate.Struct(payload)
	return err
}
