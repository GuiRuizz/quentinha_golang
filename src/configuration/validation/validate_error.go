package validation

import (
	"encoding/json"
	"errors"
	"quentinha_golang/src/configuration/rest_err"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	transL   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transL, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transL)

	}
}
func ValidateError(validationErr error, entityName string) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var validationErrors validator.ValidationErrors

	// Erro de tipo inválido no JSON (ex: string enviado para int)
	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	}

	// Erros de validação do validator
	if errors.As(validationErr, &validationErrors) {

		errorCauses := make([]rest_err.Causes, 0)

		for _, e := range validationErrors {
			cause := rest_err.Causes{
				Message: e.Translate(transL),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValidationError(
			"Some " + entityName + " fields are invalid",
			errorCauses,
		)
	}

	// Erro genérico de binding/conversão
	return rest_err.NewBadRequestError(
		"Error trying to convert " + entityName + " fields",
	)
}

