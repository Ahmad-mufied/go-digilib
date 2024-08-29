package utils

import "github.com/labstack/echo/v4"

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

func NewAPIError(code int, msg, detail string) *APIError {
	// If the detail is empty, return nil for the Detail field
	if detail == "" {
		return &APIError{Code: code, Message: msg}
	}
	return &APIError{Code: code, Message: msg, Detail: detail}
}

func HandleError(c echo.Context, err *APIError, detail ...string) error {
	// Check if a detail was provided
	if len(detail) > 0 && detail[0] != "" {
		err.Detail = detail[0]
	}
	return c.JSON(err.Code, err)
}
