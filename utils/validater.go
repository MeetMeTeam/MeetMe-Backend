package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type ValidateResponse struct {
	Message interface{} `json:"message"`
}

func CustomValidator(request interface{}) []string {
	validate := validator.New()

	var reasonErr []string

	err := validate.Struct(request)
	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {

			message := ""
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required.",
					strings.ToUpper(err.Field()))
			case "email":
				message = fmt.Sprintf("%s is not email format",
					strings.ToUpper(err.Field()))
			case "date":
				message = fmt.Sprintf("%s is valid date format",
					strings.ToUpper(err.Field()))
			case "number":
				message = fmt.Sprintf("%sis number required.",
					strings.ToUpper(err.Field()))
			case "phone":
				message = fmt.Sprintf("%s is not phone format.",
					strings.ToUpper(err.Field()))
			}

			reasonErr = append(reasonErr, message)
			// fmt.Println(err.Namespace())
			// fmt.Println(err.Field())
			// fmt.Println(err.StructNamespace())
			// fmt.Println(err.StructField())
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())
			// fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish

	}

	return reasonErr
}
