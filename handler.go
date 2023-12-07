package middleware

import (
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ScopeUnwrap() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		if !ok {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				map[string]string{"message": fmt.Sprintf("Failed to get validated JWT claims. %s", claims)},
			)
			return
		}
		scopes := strings.Split(claims.CustomClaims.(*CustomAuthClaims).Scope, " ")
		for _, v := range scopes {
			fmt.Println("setting scope in context", v)
			c.Set(v, true)
		}
		c.Next()
	}
}
