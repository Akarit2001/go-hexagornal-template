// Package apperr defines application-specific error codes and structures for consistent error handling.
package apperr

import "fmt"

// error codes are grouped by category and follow a specific range.
const (
	// Validation errors (1000–1999)
	ErrInvalidInput       = 1000
	ErrMissingField       = 1001
	ErrInvalidEmailFormat = 1002

	// Auth errors (2000–2999)
	ErrInvalidCredentials = 2000
	ErrTokenExpired       = 2001
	ErrUnauthorized       = 2002

	// Authorization (3000–3999)
	ErrAccessDenied = 3000

	// User errors (4000–4999)
	ErrUserNotFound      = 4000
	ErrUserAlreadyExists = 4001

	// Profile errors (5000–5999)
	ErrProfileNotFound    = 5000
	ErrInvalidProfileData = 5001

	// Database errors (6000–6999)
	ErrDBConnection   = 6000
	ErrDBInsertFailed = 6001
	ErrDBQueryFailed  = 6002

	// External service errors (7000–7999)
	ErrKafkaPublishFailed = 7000
	ErrHTTPCallFailed     = 7001

	// Internal (panic, unknown) errors (8000–8999)
	ErrInternalServer = 8000
)

// AppError defines a structured application error.
type AppError struct {
	Code       int    // app-specific code (e.g. "USER_NOT_FOUND")
	Message    string // user-facing message
	Err        error  // underlying error
	StatusCode int    // for HTTP responses (optional)
}

// Error satisfies the error interface.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%d: %s | %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// Unwrap returns the wrapped error (for errors.Is/As).
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new AppError.
func NewAppError(code int, message string, statusCode int, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}

func NewValidationError(message string, err error) *AppError {
	return NewAppError(1000, message, 400, err)
}

func NewNotFoundError(message string, err error) *AppError {
	return NewAppError(4000, message, 404, err)
}
