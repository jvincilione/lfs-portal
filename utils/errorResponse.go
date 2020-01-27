package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotFoundError() gin.H {
	return gin.H{
		"status":  http.StatusNotFound,
		"message": "Resource with the provided ID does not exist.",
	}
}

func GetBadRequestError(err error) gin.H {
	return gin.H{
		"status":  http.StatusBadRequest,
		"message": fmt.Sprintf("%v", err),
	}
}

func GetGenericError(err error) gin.H {
	return gin.H{
		"status":  http.StatusInternalServerError,
		"message": fmt.Sprintf("Unexpected Server Error: %v", err),
	}
}

func GetUnprocessableError(err error) gin.H {
	return gin.H{
		"status":  http.StatusUnprocessableEntity,
		"message": fmt.Sprintf("%v", err),
	}
}

func GetConflictError(err error) gin.H {
	return gin.H{
		"status":  http.StatusConflict,
		"message": fmt.Sprintf("%v", err),
	}
}

func GetUnauthorizedError() gin.H {
	return gin.H{
		"status":  http.StatusUnauthorized,
		"message": "You are not authorized",
	}
}
