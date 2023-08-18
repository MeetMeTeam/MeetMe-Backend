package utils

import (
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

			reasonErr = append(reasonErr, err.Field()+" is "+err.Tag()+".")
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
