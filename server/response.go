package server

import (
	// Standard lib
	"net/http"

	// Internal
	"github.com/marksost/img/config"

	// Third-party
	"github.com/kataras/iris"
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

// JSON outputs a JSON response via the server for a number of scenarios:
// If the environment the app is runnning is ~not~ production
// If a non-error (i.e. 200) response is detected
// If a "debug" param is passed with the request
// Otherwise, an empty text response with the proper code is output
func JSON(c *iris.Context, resp *Response) {
	// Check if this is not a production environment, or a debug flag was enabled,
	// or the status code is a non-error
	if !config.GetInstance().IsProduction() ||
		c.URLParam(DEBUG_PARAM) == "true" ||
		resp.Code == http.StatusOK {
		// Output JSON
		c.JSON(resp.Code, resp)
		return
	}

	// Output empty response as text
	c.Text(resp.Code, "")
}
