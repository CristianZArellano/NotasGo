package utils

import (
	"net/http"
	"notasGo/models"

	"github.com/gin-gonic/gin"
)

// Success responses
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, models.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error responses
func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	response := models.ErrorResponse{
		Success: false,
		Message: message,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(statusCode, response)
}

// Common error responses
func BadRequestError(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusBadRequest, message, err)
}

func NotFoundError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

func InternalServerError(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusInternalServerError, message, err)
}

func UnauthorizedError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

func ConflictError(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusConflict, message, err)
}