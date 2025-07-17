package utils

import (
	"fmt"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateStruct(data any) map[string]string {
	errors := make(map[string]string)

	err := validate.Struct(data)
	if valErrs, ok := err.(validator.ValidationErrors); ok {
		for _, v := range valErrs {
			errors[strings.ToLower(v.Field())] = getErrorMessage(v)
		}
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "email":
		return fmt.Sprintf("%s is not a valid email", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param())
	case "gte":
		return fmt.Sprintf("%s must be %s or greater", err.Field(), err.Param())

	default:
		return ""
	}
}
