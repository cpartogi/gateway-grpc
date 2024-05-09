package utils

import (
	"gateway-grpc/schema"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Success Response
func SuccessResponse(ctx echo.Context, message string, data interface{}) error {

	responseData := schema.Base{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	return ctx.JSON(http.StatusOK, responseData)
}

// Error Response
func ErrorResponse(ctx echo.Context, err error, data interface{}) error {
	statusCode, err := errorType(err)
	switch statusCode {
	case http.StatusConflict:
		return ErrorConflictResponse(ctx, err, data)
	case http.StatusBadRequest:
		return ErrorBadRequest(ctx, err, data)
	case http.StatusNotFound:
		return ErrorNotFound(ctx, err, data)
	case http.StatusForbidden:
		return ErrorForbidden(ctx, err, data)
	}
	return ErrorInternalServerResponse(ctx, err, data)
}

// ErrorConflictResponse returns
func ErrorConflictResponse(ctx echo.Context, err error, data interface{}) error {

	responseData := schema.Base{
		Status:     "Conflict",
		StatusCode: http.StatusConflict,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	return ctx.JSON(http.StatusConflict, responseData)
}

// ErrorInternalServerResponse returns
func ErrorInternalServerResponse(ctx echo.Context, err error, data interface{}) error {

	responseData := schema.Base{
		Status:     "internal server error",
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	return ctx.JSON(http.StatusInternalServerError, responseData)
}

// ErrorBadRequest returns
func ErrorBadRequest(ctx echo.Context, err error, data interface{}) error {
	responseData := schema.Base{
		Status:     "Bad Request",
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// ErrorNotFound returns
func ErrorNotFound(ctx echo.Context, err error, data interface{}) error {
	responseData := schema.Base{
		Status:     "Not found",
		StatusCode: http.StatusNotFound,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	return ctx.JSON(http.StatusNotFound, responseData)
}

// ErrorForbidden returns
func ErrorForbidden(ctx echo.Context, err error, data interface{}) error {
	responseData := schema.Base{
		Status:     "Forbidden",
		StatusCode: http.StatusForbidden,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	return ctx.JSON(http.StatusForbidden, responseData)
}
