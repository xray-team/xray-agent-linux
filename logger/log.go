package logger

import (
	"fmt"

	"github.com/go-playground/validator"
)

func LogValidationError(err error) {
	switch err.(type) {
	case validator.ValidationErrors:
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			Log.Error.Printf(Message, TagConfig, fmt.Sprintf("invalid field: '%s', invalid value: '%v', tag: '%s', param: '%s'", e.Namespace(), e.Value(), e.Tag(), e.Param()))
		}
	default:
		Log.Error.Printf(Message, TagConfig, err.Error())
	}
}
