package server

import (
	// Standard lib
	"net/http"
)

// Response is a struct defining the default shape of JSON responses this application
// should use when sending responses to HTTP requests
type Response struct {
	// HTTP status code
	Code int `json:"code"`
	// HTTP status message
	Message string `json:"message"`
	// Output data for the request
	Data interface{} `json:"data"`
}

var (
	// Common statuc code responses
	BadRequestResponse = &Response{ // 400
		Code:    http.StatusBadRequest,
		Message: http.StatusText(http.StatusBadRequest),
	}
	ForbiddenResponse = &Response{ // 403
		Code:    http.StatusForbidden,
		Message: http.StatusText(http.StatusForbidden),
	}
	MethodNotAllowedResponse = &Response{ // 405
		Code:    http.StatusMethodNotAllowed,
		Message: http.StatusText(http.StatusMethodNotAllowed),
	}
	MovedPermanentlyResponse = &Response{ // 301
		Code:    http.StatusMovedPermanently,
		Message: http.StatusText(http.StatusMovedPermanently),
	}
	NotFoundResponse = &Response{ // 404
		Code:    http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
	}
	NotImplementedResponse = &Response{ // 501
		Code:    http.StatusNotImplemented,
		Message: http.StatusText(http.StatusNotImplemented),
	}
	OKResponse = &Response{ // 200
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}
	ServerErrorResponse = &Response{ // 500
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	}
	UnauthorizedResponse = &Response{ // 401
		Code:    http.StatusUnauthorized,
		Message: http.StatusText(http.StatusUnauthorized),
	}
)
