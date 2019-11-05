package yaml

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

var (
	languages = []string{"none", "go", "ts", "js", "php", "python", "java", "cpp", "ruby"}
	syntaxes  = []string{"v1"}
)

func Validate(t Definition) error {
	validate := validator.New()

	validate.RegisterValidation("language", validateLanguage)
	validate.RegisterValidation("syntax", validateSyntax)

	err := validate.Struct(t)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Language":
				return errors.New("Malformed config, " + err.StructNamespace() + ": " + fmt.Sprintf("%v", err.Value()) + " is not a valid language")
			case "Syntax":
				return errors.New("Malformed config, " + err.StructNamespace() + ": " + fmt.Sprintf("%v", err.Value()) + " is not a valid syntax version")
			default:
				return errors.New("Malformed config, " + err.StructNamespace() + ": " + strings.ToLower(fmt.Sprintf("%v", err.StructField())) + " is required")
			}
		}
	}

	return err
}

func validateLanguage(fl validator.FieldLevel) bool {
	return contains(languages, fl.Field().String())
}

func validateSyntax(fl validator.FieldLevel) bool {
	return contains(syntaxes, fl.Field().String())
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
