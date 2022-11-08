package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"os"
)

func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func ErrorMessage(c *gin.Context, err error, statusCode int, message string) {
	if err == nil {
		log.Println("Error: " + message)
	} else {
		log.Println("Error: " + errors.Wrap(err, message).Error())
	}
	c.JSON(statusCode, gin.H{
		"message": message,
		"code":    statusCode,
	})
	c.Abort()
}
