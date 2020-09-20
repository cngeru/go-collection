package models

import "net/http"

type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func BadRequestError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NotFoundError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}
