package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Implement the error interface by adding the Error() method.
func (fe *FieldError) Error() string {
	return fmt.Sprintf("Field: %s, Message: %s", fe.Field, fe.Message)
}

func NewFieldError(field, message string) *FieldError {
	return &FieldError{
		Field:   field,
		Message: message,
	}
}
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process the request

		// Check if any errors were encountered
		if len(c.Errors) > 0 {
			// Get the last error
			err := c.Errors.Last()
			var statusCode int
			var details interface{} // Store additional details

			// Check if the error has custom metadata
			if fieldErr, ok := err.Err.(*FieldError); ok {
				// This is a custom error with field metadata
				details = map[string]interface{}{
					"field":   fieldErr.Field,
					"message": fieldErr.Message,
				}
			} else {
				// Default error details
				details = nil
			}

			// Set appropriate HTTP status code based on the error type
			switch err.Type {
			case gin.ErrorTypePublic:
				statusCode = http.StatusBadRequest // Example: Public errors
				details = map[string]interface{}{
					"hint": "Ensure the request body is valid.",
				}
			case gin.ErrorTypeBind:
				statusCode = http.StatusUnprocessableEntity // Example: Binding/Validation errors
				details = map[string]interface{}{
					"field": "email", // Example: Provide the field causing the error
					"error": err.Error(),
				}
			case gin.ErrorTypeRender:
				statusCode = http.StatusInternalServerError // Example: Rendering errors
				details = map[string]interface{}{
					"rendering_error": "An issue occurred while rendering the response.",
				}
			case gin.ErrorTypePrivate:
				statusCode = http.StatusInternalServerError // Example: Private errors
				details = map[string]interface{}{
					"error_details": "Internal server error occurred.",
				}
			case gin.ErrorTypeAny:
				statusCode = http.StatusInternalServerError // Example: Any other type of errors
				details = map[string]interface{}{
					"info": "An unknown error occurred.",
				}
			default:
				statusCode = http.StatusInternalServerError // Default: Internal server error
				details = map[string]interface{}{
					"info": "Unexpected error.",
				}
			}

			// Create the error response with additional details
			errorResponse := ErrorResponse{
				Code: statusCode,
			}
			if err.Err != nil {
				errorResponse.Message = err.Err.Error()
			}
			if details != nil {
				errorResponse.Details = details
			}

			// Respond with JSON
			c.JSON(statusCode, errorResponse)
			c.Abort()
		}
	}
}
