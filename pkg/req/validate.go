package req

import "github.com/go-playground/validator/v10"

func IsValid[T any](obj *T) error {
	validate := validator.New()
	return validate.Struct(obj)
}
