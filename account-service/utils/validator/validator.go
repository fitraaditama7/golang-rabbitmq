package validator

import (
	"account-service/utils/logger"
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func validateMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error() // default error
}

func Validate(a interface{}) error {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	err := validate.Struct(a)
	if err != nil {
		l := logger.Log()
		l.Error().Err(err).Msg(err.Error())
		validationErr := err.(validator.ValidationErrors)[0]
		return errors.New(validateMessage(validationErr))
	}
	return nil
}
