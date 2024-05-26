package handle_erro

import "github.com/go-playground/validator/v10"

func ValidationError(err error) map[string]string {
	var errors map[string]string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors = make(map[string]string)
		for _, validationErr := range validationErrors {
			var errorMsg string
			switch validationErr.Tag() {
			case "required":
				errorMsg = validationErr.Field() + " is required."
			case "min":
				errorMsg = validationErr.Field() + " must be at least " + validationErr.Param() + "characters."
			case "max":
				errorMsg = validationErr.Field() + " must be at most " + validationErr.Param() + "characters."
			case "email":
				errorMsg = validationErr.Field() + " must be a valid E-mail."
			default:
				errorMsg = validationErr.Field() + " is invalid."
			}
			errors["error"] = errorMsg
		}
	} else {
		errors = map[string]string{"error": err.Error()}
	}
	return errors
}
