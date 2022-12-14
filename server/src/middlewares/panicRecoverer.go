package middlewares

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PanicRecoverer() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
