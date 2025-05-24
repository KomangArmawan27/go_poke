package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BaseResponse formats error messages consistently
type BaseResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ValidationErrorResponse(c *gin.Context, err error) {
	var errors []string
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			switch e.Tag() {
			case "required":
				errors = append(errors, fmt.Sprintf("%s is required", e.Field()))
			case "min":
				errors = append(errors, fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param()))
			default:
				errors = append(errors, e.Field()+": "+e.Tag()+" "+e.Param())
			}
		}
		Response(c, 400, false, "Validation failed", gin.H{
			"validation_errors": errors,
		})
	} else {
		Response(c, 400, false, err.Error(), nil)
	}
}

func Response(c *gin.Context, code int, success bool, message string, data interface{}) {
	response := BaseResponse{
		Success: success,
		Code:    code,
		Message: message,
		Data:    data,
	}

	c.JSON(code, response)
}
