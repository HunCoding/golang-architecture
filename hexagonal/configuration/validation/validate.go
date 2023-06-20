package validation

import (
	"encoding/json"
	"errors"
	"hexagonal/configuration/rest_errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enTranslator := en.New()
		unt := ut.New(enTranslator, enTranslator)
		transl, _ = unt.GetTranslator("en")
		err := en_translation.RegisterDefaultTranslations(val, transl)
		if err != nil {
			return
		}
	}
}

func ValidateUserError(
	validation_err error,
) *rest_errors.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_errors.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_errors.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_errors.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_errors.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_errors.NewBadRequestError("Error trying to convert fields")
	}
}
