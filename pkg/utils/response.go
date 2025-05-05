package utils

import (
	"github.com/gin-gonic/gin"
)

// Response is the standard API response structure
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// RespondWithSuccess sends a successful response with data
func RespondWithSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Data:    data,
	})
}

// RespondWithError sends an error response
func RespondWithError(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, Response{
		Success: false,
		Error:   err,
	})
}
