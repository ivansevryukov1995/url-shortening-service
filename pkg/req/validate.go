package req

import "github.com/go-playground/validator/v10"

func IsValid[T any](req T) error {
	validate := validator.New()
	err := validate.Struct(req)
	return err

}
