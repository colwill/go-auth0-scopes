package gin_http

import (
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	middleware "go-auth0-scopes"
	"net/http"
	"net/url"
	"time"
)

var router = gin.Default()

// These must be set to valid values defined in your auth0 account

const auth0Project = "AUTH0_PROJECT"   // e.g. myproject
const auth0Location = "AUTH0_LOCATION" // e.g. eu
const auth0Audience = "AUTH0_AUDIENCE" // e.g. api

func main() {
	gin.SetMode(gin.DebugMode)
	issuerUrl, _ := url.Parse(fmt.Sprintf("https://%s.%s.auth0.com", auth0Project, auth0Location))
	issuerBaseUrl, _ := url.Parse(fmt.Sprintf("https://%s.%s.auth0.com/", auth0Project, auth0Location))
	customClaims := func() validator.CustomClaims {
		return &middleware.CustomAuthClaims{}
	}
	provider := jwks.NewCachingProvider(issuerUrl, 5*time.Minute)
	jwtValidator, _ := validator.New(provider.KeyFunc,
		validator.RS256,
		issuerBaseUrl.String(),
		[]string{auth0Audience},
		validator.WithCustomClaims(customClaims),
	)
	jwtMiddleware := jwtmiddleware.New(jwtValidator.ValidateToken)
	engine := gin.Engine{}
	err := engine.SetTrustedProxies([]string{"0.0.0.0/0", "::/0"})
	if err != nil {
		return
	}
	router.Use(adapter.Wrap(jwtMiddleware.CheckJWT))
	router.Use(middleware.ScopeUnwrap())

	// unprotected route

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "healthy"})
	})

	// protected route using
	// middleware.Authorise("read:protected")
	// the read:protected scope must be assigned to the clientId calling this endpoint
	// if the clientId does not have this scope assigned, the result will be a 401 Unauthorized

	router.GET("/protected", middleware.Authorise("read:protected"), func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "this is a protected route!"})
	})

	err = router.Run()
	if err != nil {
		return
	}
}
