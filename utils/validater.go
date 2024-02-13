package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"meetme/be/errs"
	"os"
	"strings"

	"github.com/go-playground/validator"
)

type ValidateResponse struct {
	Message interface{} `json:"message"`
}

type Claims struct {
	Email     string  `json:"email"`
	IsRefresh bool    `json:"isRefresh"`
	IsAdmin   bool    `json:"isAdmin"`
	ExpiredAt float64 `json:"expiredAt"`
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
			case "unique":
				message = fmt.Sprintf("%s is already.",
					strings.ToUpper(err.Field()))
			}

			reasonErr = append(reasonErr, message)

		}

		// from here you can create your own error messages in whatever language you wish

	}

	return reasonErr
}

func IsTokenValid(authHeader string) (*Claims, error) {
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		// ถอดรหัส Token โดยตัด "Bearer " ออก
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// แกะ Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// ในกรณีที่ใช้การเก็บคีย์เป็นสาธารณะและส่วนตัว
			// คุณสามารถส่งคีย์ในนี้ แต่ควรใช้แนวทางที่ปลอดภัยกว่าในบริการจริง
			return []byte(os.Getenv("APP_SECRET")), nil
		})

		if err != nil {
			return nil, errs.NewUnauthorizedError("Invalid Token: " + err.Error())
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// ตรวจสอบและเข้าถึงข้อมูลที่คุณต้องการจาก claims

			claimsRes := Claims{
				Email:     claims["email"].(string),
				IsRefresh: claims["isRefresh"].(bool),
				ExpiredAt: claims["exp"].(float64),
			}

			return &claimsRes, nil
		} else {
			return nil, errs.NewUnauthorizedError("Invalid Token: " + err.Error())
		}

	} else {
		return nil, errs.NewUnauthorizedError("Invalid or missing Bearer Token")
	}
}
