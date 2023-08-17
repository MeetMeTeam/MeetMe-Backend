package utils

type DataResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// type CustomValidator struct {
// 	Validator *validator.Validate
// }

// func (cv *CustomValidator) Validate(i interface{}) error {
// 	if err := cv.Validator.Struct(i); err != nil {
// 		// Optionally, you could return the error to give each route more control over the status code
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return nil
// }
