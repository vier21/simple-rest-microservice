package utils

import "github.com/go-playground/validator/v10"

func ValidateStruct(v *validator.Validate, obj any) error {
	return v.Struct(obj)
}