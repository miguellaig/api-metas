package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{err.Error()}
	}

	var messages []string
	for _, e := range validationErrors {
		var msg string
		switch e.Tag() {
		case "required":
			msg = fmt.Sprintf("O campo '%s' é obrigatório", e.Field())
		case "min":
			msg = fmt.Sprintf("O campo '%s' deve ter no mínimo %s caracteres", e.Field(), e.Param())
		case "max":
			msg = fmt.Sprintf("O campo '%s' deve ter no máximo %s caracteres", e.Field(), e.Param())
		case "oneof":
			msg = fmt.Sprintf("O campo '%s' deve ser um dos valores permitidos", e.Field())
		default:
			msg = fmt.Sprintf("O campo '%s' falhou na validação '%s'", e.Field(), e.Tag())
		}
		messages = append(messages, msg)
	}
	return messages
}
