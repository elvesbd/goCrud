package validation

import (
	"encoding/json"
	"errors"

	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *restErr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return restErr.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		errorsCause := []restErr.Causes{}

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := restErr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCause = append(errorsCause, cause)
		}
		return restErr.NewBadRequestValidationError("Some fields are invalid", errorsCause)
	} else {
		return restErr.NewBadRequestError("Error trying to convert fields")
	}
}
