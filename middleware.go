package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorise(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Keys[scope] == nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				map[string]string{"message": fmt.Sprintf("Scope required to perform action: %s", scope)},
			)
			return
		} else {
			c.Next()
		}
	}
}
