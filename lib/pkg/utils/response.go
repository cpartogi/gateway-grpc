package utils

import (
	"errors"
	"gateway-grpc/schema"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	statusCode, _ := status.FromError(err)
	switch statusCode.Code() {
	case codes.AlreadyExists:
		return ErrorConflictResponse(ctx, errors.New("already exist"), data)
	case codes.InvalidArgument:
		return ErrorBadRequest(ctx, errors.New(statusCode.Message()), data)
	case codes.NotFound:
		return ErrorNotFound(ctx, errors.New("data not found"), data)
	case codes.PermissionDenied:
		return ErrorForbidden(ctx, errors.New("permission denied"), data)
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
